package middlewares

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/gofiber/fiber"
	"ms.api/models"
	"ms.api/utils"
	"net"
	"net/http"
	"regexp"
)

var ProtectedCalls []string

func init() {
	// Mount all protected endpoints here.
	ProtectedCalls = calls
}

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
	for _, protectedCall := range ProtectedCalls {
		// Match things like this `QueryOrMutationName(` or `QueryOrMutationName{` even if there is a whitespace between `...Name {` or `...Name (` specific target.
		match := fmt.Sprintf(`(?m)%s\s*[\(\{]`, protectedCall)
		r := regexp.MustCompile(match)
		if found := r.FindAllString(query, -1); len(found) > 0 {
			return protectedCall, true
		}
	}
	return "", false
}

func ProtectedMiddleware(c *fiber.Ctx) {
	query := string(c.Fasthttp.Request.Body())
	if _, isProtected := ValidateAPICall(query); isProtected {
		AuthenticatedUser, ok := c.Fasthttp.Value(AuthenticatedUserContextKey).(utils.JSON)
		if !ok || AuthenticatedUser == nil {
			result := models.Result{
				Success: false,
				Message: "Sorry, you must be authenticated/logged in to continue.",
				//ReturnStatus: models.ReturnStatusAuthenticationError,
			}
			c.Status(http.StatusUnauthorized)
			_ = c.JSON(result)
			return
		}
	}

	c.Next()
	return
}
