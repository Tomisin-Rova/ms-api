package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"ms.api/fakes"
	"ms.api/protos/pb/onboardingService"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHttpHandler_VerifyMagicLinkHandler(t *testing.T) {
	svc := fakes.NewFakeOnBoardingClient(&onboardingService.SuccessResponse{Message: "success"},
		nil, nil, nil)

	handler := New(svc, logrus.StandardLogger())
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	u, err := url.Parse("http://localhost.app/verify_email?email=foo@bar.io&verificationToken=token")
	assert.Nil(t, err)
	r.URL = u

	handler.VerifyMagicLinkHandler(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "<h1>email has been successfully verified</h1>", w.Body.String())
}

func TestHttpHandler_VerifyMagicLinkHandler_BadEmail(t *testing.T) {
	svc := fakes.NewFakeOnBoardingClient(&onboardingService.SuccessResponse{Message: "success"},
		nil, nil, nil)

	handler := New(svc, logrus.StandardLogger())
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	u, err := url.Parse("http://localhost.app/verify_email?email=foo@bar&verificationToken=token")
	assert.Nil(t, err)
	r.URL = u

	handler.VerifyMagicLinkHandler(w, r)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "<h1>invalid email address</h1>", w.Body.String())
}
