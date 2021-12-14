package handlers

import (
	"go.uber.org/zap"
)

type HttpHandler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *HttpHandler {
	return &HttpHandler{logger: logger}
}
