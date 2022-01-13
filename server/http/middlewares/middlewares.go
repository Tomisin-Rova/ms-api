package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/roava/zebra/models"
	"google.golang.org/grpc/metadata"
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/types"

	coreMiddleware "github.com/roava/zebra/middleware"
	"go.uber.org/zap"
)

const (
	AuthenticatedUserContextKey = "AuthenticatedUser"
	Bearer                      = "Bearer"
)

type AuthMiddleware struct {
	authService auth.AuthServiceClient
	logger      *zap.Logger
}

func NewAuthMiddleware(service auth.AuthServiceClient, logger *zap.Logger) *AuthMiddleware {
	return &AuthMiddleware{authService: service, logger: logger}
}

func (mw *AuthMiddleware) ValidateToken(token string) (*models.JWTClaims, error) {
	// Execute RPC call
	response, err := mw.authService.ValidateToken(context.Background(), &auth.ValidateTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	// Build response
	jwtClaims := models.JWTClaims{
		ID:       response.Claims.Id,
		Email:    response.Claims.Email,
		DeviceID: response.Claims.DeviceId,
	}
	switch response.Claims.ClientType {
	case types.JWTClaims_APP:
		jwtClaims.Client = models.APP
	case types.JWTClaims_DASHBOARD:
		jwtClaims.Client = models.DASHBOARD
	}

	return &jwtClaims, nil
}

// GetClaimsFromCtx returns claims from an authenticated user
func GetClaimsFromCtx(ctx context.Context) (*models.JWTClaims, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return nil, errors.New("unable to parse authenticated user")
	}
	jsonClaims := md.Get(AuthenticatedUserContextKey)
	if len(jsonClaims) == 0 {
		return nil, errors.New("fail decode authenticated user claims")
	}

	var claims models.JWTClaims
	err := json.Unmarshal([]byte(jsonClaims[0]), &claims)
	if err != nil {
		return nil, err
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

		ctx, err := coreMiddleware.PutClaimsOnContext(r.Context(), &models.JWTClaims{
			Client:   resp.Client,
			ID:       resp.ID,
			Email:    resp.Email,
			DeviceID: resp.DeviceID,
		})
		if err != nil {
			mw.logger.Info(fmt.Sprintf("failed to marshal claims: %v", err),
				zap.String("token", token),
			)
			next.ServeHTTP(w, r)
			return
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
