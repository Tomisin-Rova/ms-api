package middlewares

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/roava/zebra/middleware"
	"github.com/roava/zebra/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"google.golang.org/grpc/metadata"
	"ms.api/mocks"
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/types"
)

func Test_GetClaimsFromCtx(t *testing.T) {
	const (
		failNotOutgoingContext = iota
		failToDecodeAuthenticatedUserClaims
		failToUnmarshalClaims
		success
	)

	successCtx, _ := middleware.PutClaimsOnContext(context.Background(), &models.JWTClaims{ID: "user-id"})

	testCases := []struct {
		name     string
		arg      context.Context
		testType int
	}{
		{
			name:     "fail to parse authenticated user",
			arg:      metadata.NewIncomingContext(context.Background(), nil),
			testType: failNotOutgoingContext,
		},
		{
			name:     "fail to decode user claims",
			arg:      metadata.NewOutgoingContext(context.Background(), nil),
			testType: failToDecodeAuthenticatedUserClaims,
		},
		{
			name:     "fail to unmarshall claims",
			arg:      metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{middleware.AuthenticatedUserMetadataKey: ""})),
			testType: failToUnmarshalClaims,
		},
		{
			name:     "success",
			arg:      successCtx,
			testType: success,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			switch testCase.testType {
			case failNotOutgoingContext:
				claims, err := GetClaimsFromCtx(testCase.arg)
				assert.Error(t, err)
				assert.Equal(t, err.Error(), "unable to parse authenticated user")
				assert.Nil(t, claims)
			case failToDecodeAuthenticatedUserClaims:
				claims, err := GetClaimsFromCtx(testCase.arg)
				assert.Error(t, err)
				assert.Equal(t, err.Error(), "fail decode authenticated user claims")
				assert.Nil(t, claims)
			case failToUnmarshalClaims:
				claims, err := GetClaimsFromCtx(testCase.arg)
				assert.Error(t, err)
				assert.Equal(t, err.Error(), "fail to unmarshall claims")
				assert.Nil(t, claims)
			case success:
				claims, err := GetClaimsFromCtx(testCase.arg)
				assert.NoError(t, err)
				assert.NotNil(t, claims)
				assert.Equal(t, "user-id", claims.ID)
			}
		})
	}
}

type Handler struct {
	T *testing.T
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	claims, err := GetClaimsFromCtx(r.Context())
	assert.NoError(h.T, err)
	assert.NotNil(h.T, claims)
}

func Test_Middleware(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockAuthService := mocks.NewMockAuthServiceClient(controller)
	req := http.Request{
		Header: map[string][]string{
			"Authorization": {
				"Bearer somejwtclaim",
			},
		},
	}

	mockAuthService.EXPECT().ValidateToken(gomock.Any(), gomock.Any()).Return(&auth.ValidateTokenResponse{
		IsValid: true,
		Claims: &types.JWTClaims{
			Id:       "some-id",
			Email:    "user@email.org",
			DeviceId: "some-device-id",
		},
	}, nil).Times(1)

	service := NewAuthMiddleware(mockAuthService, zaptest.NewLogger(t))
	handler := service.Middleware(&Handler{t})
	handler.ServeHTTP(httptest.NewRecorder(), &req)
}
