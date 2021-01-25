package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"ms.api/protos/pb/authService"
	"net/http"
	"strings"
)

type ctxKey struct {
	Name string
}

var (
	AuthenticatedUserContextKey = &ctxKey{Name: "AuthenticatedUser"}
)

const (
	Bearer string = "Bearer"
)

type AuthMiddleware struct {
	authService authService.AuthServiceClient
	logger      *zap.Logger
}

func NewAuthMiddleware(service authService.AuthServiceClient, logger *zap.Logger) *AuthMiddleware {
	return &AuthMiddleware{authService: service, logger: logger}
}

func (mw *AuthMiddleware) ValidateToken(token string) (*authService.ValidateTokenResponse, error) {
	resp, err := mw.authService.ValidateToken(context.Background(),
		&authService.ValidateTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TODO: here user should be the direct type of protos.Person from the auth or person service.
func GetAuthenticatedUser(ctx context.Context) (*models.Claims, error) {
	claims, ok := ctx.Value(AuthenticatedUserContextKey).(models.Claims)
	if !ok {
		return nil, errors.New("unable to parse authenticated user")
	}
	return &claims, nil
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

		resp, err := mw.ValidateToken(token)
		if err != nil {
			mw.logger.Info(fmt.Sprintf("failed to validate token: %v", err),
				zap.String("token", token),
			)
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), AuthenticatedUserContextKey, models.Claims{
			PersonId:   resp.PersonId,
			IdentityId: resp.IdentityId,
			DeviceId:   resp.DeviceId,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
