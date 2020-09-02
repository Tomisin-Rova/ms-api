package middlewares

import (
	"context"
	"ms.api/types"
	"ms.api/utils"
	"net/http"
	"strings"

	"ms.api/libs/sessions"
)

const (
	AuthenticatedUserContextKey = "AuthenticatedUser"
	TokenContextKey             = "Token"
	//AccountTypeContextKey       = "AccountType"
	Bearer = "BEARER"
)

func handleSessionByToken(token string) (AuthenticatedUser utils.JSON, result *types.Result) {
	AuthenticatedUser = make(utils.JSON)

	session, err := sessions.GetSessionByToken(token)
	if err != nil {
		result := &types.Result{}
		result.Success = false
		result.Message = "Sorry, your session has expired, Please login to continue. "
		//result.ReturnStatus = types.ReturnStatusTokenExpired
		return nil, result
	}

	println(session.AccountId)
	// TODO: Come here and apply an actual session thing for the connected client to carry out any sensitive business.
	//switch accountType.String() {

	//case types.AccountTypesAdministrator.String():
	//	AuthenticatedUser[types.AccountTypesAdministrator.String()], _ = Administrators.ViewAdmin(_id)
	//	break
	//case types.AccountTypesScheduler.String():
	//	AuthenticatedUser[types.AccountTypesScheduler.String()], _ = Schedulers.ViewScheduler(_id)
	//	break
	//case types.AccountTypesSigner.String():
	//	AuthenticatedUser[types.AccountTypesSigner.String()], _ = Signers.ViewSigner(_id)
	//	break
	//}
	//_ = sessions.ExtendSession(token)
	return AuthenticatedUser, nil
}
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
	return sessions.DestroySession(token)
}
