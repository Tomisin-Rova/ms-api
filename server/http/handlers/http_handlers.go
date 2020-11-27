package handlers

import (
	"context"
	"github.com/sirupsen/logrus"
	rerrors "ms.api/libs/errors"
	emailvalidator "ms.api/libs/validator/email"
	"ms.api/protos/pb/onboardingService"
	"net/http"
)

type HttpHandler struct {
	onboardingService onboardingService.OnBoardingServiceClient
	logger            *logrus.Logger
}

func New(svc onboardingService.OnBoardingServiceClient, logger *logrus.Logger) *HttpHandler {
	return &HttpHandler{onboardingService: svc, logger: logger}
}

func (handler *HttpHandler) VerifyMagicLinkHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	email, verificationToken := q.Get("email"), q.Get("verificationToken")
	if err := emailvalidator.Validate(email); err != nil {
		handler.logger.WithField("email", email).
			WithError(err).
			Error("invalid email supplied")
		handler.respond(w, http.StatusBadRequest, "<h1>invalid email address</h1>")
		return
	}
	resp, err := handler.onboardingService.VerifyEmailMagicLInk(context.Background(), &onboardingService.VerifyEmailMagicLInkRequest{
		Email:             email,
		VerificationToken: verificationToken,
	})
	if err != nil {
		handler.logger.WithError(err).Error("failed to call verify magic link")
		handler.respond(w, http.StatusInternalServerError, rerrors.NewFromGrpc(err).Error())
		return
	}
	handler.logger.WithField("response", resp).Info("magic link verified")
	handler.respond(w, http.StatusOK, "<h1>email has been successfully verified</h1>")
}

func (handler *HttpHandler) respond(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Write([]byte(message))
}
