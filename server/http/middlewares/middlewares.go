package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/roava/zebra/models"
	"ms.api/protos/pb/auth"

	"go.uber.org/zap"
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
	authService auth.AuthServiceClient
	logger      *zap.Logger
}

func NewAuthMiddleware(service auth.AuthServiceClient, logger *zap.Logger) *AuthMiddleware {
	return &AuthMiddleware{authService: service, logger: logger}
}

func (mw *AuthMiddleware) ValidateToken(token string) (*models.JWTClaims, error) {
	// TODO: Implement logic once auth service is refactored .
	return &models.JWTClaims{
		Client:   models.ClientType(models.APP),
		ID:       "01fk5jmz4thmxwz8p2fx45vj6v",
		Email:    "fola@roava.app",
		DeviceID: "01f82zca7ryacseqddc8a6twte",
	}, nil
}

// TODO: here user should be the direct type of protos.Person from the auth or person service.
func GetAuthenticatedUser(ctx context.Context) (*models.JWTClaims, error) {
	claims, ok := ctx.Value(AuthenticatedUserContextKey).(models.JWTClaims)
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

		ctx := context.WithValue(r.Context(), AuthenticatedUserContextKey, models.JWTClaims{
			Client:   resp.Client,
			ID:       resp.ID,
			Email:    resp.Email,
			DeviceID: resp.DeviceID,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
