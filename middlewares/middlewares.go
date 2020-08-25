package middlewares

import (
	"context"
	"github.com/gofiber/fiber"
	"ms.api/libs/ObjectID"
	"ms.api/utils"
	"net/http"
	"strings"

	"ms.api/libs/sessions"
	"ms.api/models"
)

const (
	AuthenticatedUserContextKey = "AuthenticatedUser"
	TokenContextKey             = "Token"
	//AccountTypeContextKey       = "AccountType"
	Bearer = "BEARER"
)

func _handleSessionByToken(token string) (AuthenticatedUser utils.JSON, result *models.Result) {
	AuthenticatedUser = make(utils.JSON)

	session, err := sessions.GetSessionByToken(token)
	if err != nil {
		result := &models.Result{}
		result.Success = false
		result.Message = "Sorry, your session has expired, Please login to continue. "
		//result.ReturnStatus = models.ReturnStatusTokenExpired
		return nil, result
	}

	_, _ = ObjectID.UnmarshalID(session.AccountId)
	//switch accountType.String() {

	//case models.AccountTypesAdministrator.String():
	//	AuthenticatedUser[models.AccountTypesAdministrator.String()], _ = Administrators.ViewAdmin(_id)
	//	break
	//case models.AccountTypesScheduler.String():
	//	AuthenticatedUser[models.AccountTypesScheduler.String()], _ = Schedulers.ViewScheduler(_id)
	//	break
	//case models.AccountTypesSigner.String():
	//	AuthenticatedUser[models.AccountTypesSigner.String()], _ = Signers.ViewSigner(_id)
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

func AuthMiddleWare(c *fiber.Ctx) {
	Authorization := c.Get("Authorization")
	var token string

	if Authorization == "" {
		c.Next()
		return
	}

	if len(Authorization) > len(Bearer) && strings.ToUpper(Authorization[0:len(Bearer)]) == Bearer {
		token = Authorization[len(Bearer)+1:]
	} else {
		c.Next(nil)
		return
	}

	AuthenticatedUser, result := _handleSessionByToken(token)
	if result != nil {
		c.Status(http.StatusUnauthorized)
		_ = c.JSON(result)
		return
	}

	AuthenticatedUser[TokenContextKey] = token

	c.Fasthttp.SetUserValue(AuthenticatedUserContextKey, AuthenticatedUser)
	c.Next(nil)
	return
}

func DestroyAuthenticatedUser(ctx context.Context) error {
	_, token := GetAuthenticatedUser(ctx)
	if token == "" {
		return nil
	}
	return sessions.DestroySession(token)
}
