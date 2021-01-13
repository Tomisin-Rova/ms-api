package middlewares

import (
	"errors"
	"github.com/roava/zebra/logger"
	"github.com/stretchr/testify/assert"
	"ms.api/fakes"
	"ms.api/protos/pb/authService"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware_Middeware_InvalidJWT(t *testing.T) {
	authClient := fakes.NewFakeAuthClient(nil, nil, nil, errors.New("cannot validate token"))
	mw := NewAuthMiddleware(authClient, logger.New())

	fn := func(w http.ResponseWriter, r *http.Request) {
		personId, err := GetAuthenticatedUser(r.Context())
		assert.NotNil(t, err)
		assert.Empty(t, personId)
	}
	handler := mw.Middeware(http.HandlerFunc(fn))
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)

	r.Header.Set("Authorization", "Bearer JWT")
	handler.ServeHTTP(w, r)
}

func TestAuthMiddleware_Middeware_Success(t *testing.T) {
	authClient := fakes.NewFakeAuthClient(&authService.ValidateTokenResponse{PersonId: "personId"},
		nil, nil, nil)
	mw := NewAuthMiddleware(authClient, logger.New())

	fn := func(w http.ResponseWriter, r *http.Request) {
		personId, err := GetAuthenticatedUser(r.Context())
		assert.Nil(t, err)
		assert.Equal(t, personId, "personId")
	}
	handler := mw.Middeware(http.HandlerFunc(fn))
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)

	r.Header.Set("Authorization", "Bearer JWT")
	handler.ServeHTTP(w, r)
}

func TestAuthMiddleware_Middeware_Header_Absent(t *testing.T) {
	authClient := fakes.NewFakeAuthClient(nil,
		nil, nil, errors.New("failed to validate token"))
	mw := NewAuthMiddleware(authClient, logger.New())

	fn := func(w http.ResponseWriter, r *http.Request) {
		personId, err := GetAuthenticatedUser(r.Context())
		assert.NotNil(t, err)
		assert.Empty(t, personId)
	}
	handler := mw.Middeware(http.HandlerFunc(fn))
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)

	// r.Header.Set("Authorization", "Bearer JWT")
	handler.ServeHTTP(w, r)
}
