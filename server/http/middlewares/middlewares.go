package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	coreMiddleware "github.com/roava/zebra/middleware"
	"github.com/roava/zebra/models"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	errorvalues "ms.api/libs/errors"
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/types"
)

const (
	Bearer = "Bearer"
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
		return nil, errorvalues.Format(errorvalues.InvalidAuthentication, errors.New("unable to parse authenticated user"))
	}
	jsonClaims := md.Get(coreMiddleware.AuthenticatedUserMetadataKey)
	if len(jsonClaims) == 0 {
		return nil, errorvalues.Format(errorvalues.InvalidAuthentication, errors.New("fail decode authenticated user claims"))
	}

	var claims models.JWTClaims
	err := json.Unmarshal([]byte(jsonClaims[0]), &claims)
	if err != nil {
		return nil, errorvalues.Format(errorvalues.InvalidAuthentication, errors.New("fail to unmarshall claims"))
	}

	return &claims, nil
}

func (mw *AuthMiddleware) Middleware(next http.Handler) http.Handler {
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

		ctx, _ := coreMiddleware.PutClaimsOnContext(r.Context(), &models.JWTClaims{
			Client:   resp.Client,
			ID:       resp.ID,
			Email:    resp.Email,
			DeviceID: resp.DeviceID,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
