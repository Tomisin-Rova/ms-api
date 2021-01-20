package handlers

import (
	"context"
	rerrors "ms.api/libs/errors"
	"net/http"

	emailvalidator "ms.api/libs/validator/email"
	"ms.api/protos/pb/onboardingService"

	"go.uber.org/zap"
)

type HttpHandler struct {
	onboardingService onboardingService.OnBoardingServiceClient
	logger            *zap.Logger
}

func New(svc onboardingService.OnBoardingServiceClient, logger *zap.Logger) *HttpHandler {
	return &HttpHandler{onboardingService: svc, logger: logger}
}

func (handler *HttpHandler) VerifyMagicLinkHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	email, verificationToken := q.Get("email"), q.Get("verificationToken")
	if err := emailvalidator.Validate(email); err != nil {
		handler.logger.Error("invalid email supplied",
			zap.Error(err),
			zap.String("email", email),
		)
		handler.respond(w, http.StatusBadRequest, "<h1>invalid email address</h1>")
		return
	}
	resp, err := handler.onboardingService.VerifyEmailMagicLInk(context.Background(), &onboardingService.VerifyEmailMagicLInkRequest{
		Email:             email,
		VerificationToken: verificationToken,
	})
	if err != nil {
		handler.logger.Error("failed to call verify magic link", zap.Error(err))
		handler.respond(w, http.StatusInternalServerError, rerrors.NewFromGrpc(err).Error())
		return
	}
	handler.logger.Info("magic link verified", zap.Any("response", resp))
	handler.respond(w, http.StatusOK, "<h1>email has been successfully verified</h1>")
}

func (handler *HttpHandler) respond(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	if _, err := w.Write([]byte(message)); err != nil {
		handler.logger.Error("error occurred while sending data to client", zap.Error(err))
	}
}
