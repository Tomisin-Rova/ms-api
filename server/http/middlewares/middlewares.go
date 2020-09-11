package middlewares

import (
	"context"
	"ms.api/types"
	"ms.api/utils"
	"net/http"
	"strings"
)

const (
	AuthenticatedUserContextKey = "AuthenticatedUser"
	TokenContextKey             = "Token"
	//AccountTypeContextKey       = "AccountType"
	Bearer = "BEARER"
)

func handleSessionByToken(token string) (AuthenticatedUser utils.JSON, result *types.Result) {
	AuthenticatedUser = make(utils.JSON)
	// TODO: use token to call ms.auth to return person & identity for use by this gateway going forward.
	//AuthenticatedUser["person"], _ = authService.GetUserFromToken(token)
	return AuthenticatedUser, nil
}

// TODO: here user should be the direct type of protos.Person from the auth or person service.
func GetAuthenticatedUser(ctx context.Context) (user interface{}, token string) {
	authenticatedUser, _ := ctx.Value(AuthenticatedUserContextKey).(utils.JSON)
	user = authenticatedUser["user"]
	token = authenticatedUser[TokenContextKey].(string)
	return user, token
}

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		Authorization := r.Header.Get("Authorization")
		var token string

		if Authorization == "" {
			next.ServeHTTP(w, r)
			return
		}

		if len(Authorization) > len(Bearer) && strings.ToUpper(Authorization[0:len(Bearer)]) == Bearer {
			token = Authorization[len(Bearer)+1:]
		} else {
			next.ServeHTTP(w, r)
			return
		}

		AuthenticatedUser, result := handleSessionByToken(token)
		if result != nil {
			w.WriteHeader(http.StatusUnauthorized)
			// TODO: Write the bad result here.
			next.ServeHTTP(w, r)
			return
		}

		AuthenticatedUser[TokenContextKey] = token

		ctx := context.WithValue(r.Context(), AuthenticatedUserContextKey, AuthenticatedUser)
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}

func DestroyAuthenticatedUser(ctx context.Context) error {
	_, token := GetAuthenticatedUser(ctx)
	if token == "" {
		return nil
	}
	// TODO: Call ms.auth to logout this authenticated user via a token.
	return nil
}
