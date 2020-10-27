package middlewares

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"ms.api/types"
	"ms.api/utils"
	"net"
	"net/http"
	"regexp"
)

type ResWriter struct {
	http.ResponseWriter
	http.Hijacker
	buf *bytes.Buffer
}

// Implement the interface for WebSockets to be able to upgrade connection.
func (mrw *ResWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := mrw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the ResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

func (mrw *ResWriter) Write(p []byte) (int, error) {
	return mrw.buf.Write(p)
}

func ValidateAPICall(query string) (string, bool) {
	for _, protectedCall := range protectedCalls {
		// Match things like this `QueryOrMutationName(` or `QueryOrMutationName{` even if there is a whitespace between `...Name {` or `...Name (` specific target.
		match := fmt.Sprintf(`(?m)%s\s*[\(\{]`, protectedCall)
		r := regexp.MustCompile(match)
		if found := r.FindAllString(query, -1); len(found) > 0 {
			return protectedCall, true
		}
	}
	return "", false
}

func ProtectedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check Protected Routes here.
		s, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(s))
		query := string(s) // We can further do checks with this query body here.

		if _, isProtected := ValidateAPICall(query); isProtected {
			ctx := r.Context()
			AuthenticatedUser, ok := ctx.Value(AuthenticatedUserContextKey).(utils.JSON)
			if !ok || AuthenticatedUser == nil {
				result := types.Result{
					Success: false,
					Message: "Sorry, you must be authenticated/logged in to continue.",
					//ReturnStatus: models.ReturnStatusAuthenticationError,
				}
				w.WriteHeader(http.StatusUnauthorized)
				b, _ := json.Marshal(result)
				_, _ = w.Write(b)
				return
			}
		}

		protected := &ResWriter{
			ResponseWriter: w,
			buf:            &bytes.Buffer{},
		}

		next.ServeHTTP(protected, r)
		if _, e := io.Copy(w, protected.buf); e != nil {
			// To handle copying a dead-body.
			return
		}
	})
}
