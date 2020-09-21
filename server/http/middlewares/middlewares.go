package middlewares

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"ms.api/protos/pb/authService"
	"net/http"
	"strings"
)

const (
	AuthenticatedUserContextKey = "AuthenticatedUser"
	Bearer                      = "Bearer"
)

type AuthMiddleware struct {
	authService authService.AuthServiceClient
	logger      *logrus.Logger
}

func NewAuthMiddleware(service authService.AuthServiceClient, logger *logrus.Logger) *AuthMiddleware {
	return &AuthMiddleware{authService: service, logger: logger}
}

func (mw *AuthMiddleware) ValidateToken(token string) (string, error) {
	resp, err := mw.authService.ValidateToken(context.Background(),
		&authService.ValidateTokenRequest{Token: token})
	if err != nil {
		return "", err
	}
	return resp.PersonId, nil
}

// TODO: here user should be the direct type of protos.Person from the auth or person service.
func GetAuthenticatedUser(ctx context.Context) (string, error) {
	personId, ok := ctx.Value(AuthenticatedUserContextKey).(string)
	if !ok {
		return "", errors.New("unable to parse authenticated user")
	}
	return personId, nil
}

func (mw *AuthMiddleware) Middeware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		token := ""
		if authorization == "" {
			next.ServeHTTP(w, r)
			return
		}

		sp := strings.Split(authorization, " ")
		if len(sp) > 1 {
			token = sp[1]
		}

		if token == "" {
			mw.logger.Info("no token supplied")
			next.ServeHTTP(w, r)
			return
		}

		personId, err := mw.ValidateToken(token)
		if err != nil {
			mw.logger.WithField("token", token).Infof("failed to validate token: %v", err)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), AuthenticatedUserContextKey, personId)
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}
