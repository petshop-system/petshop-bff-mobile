package customer

import (
	"go.uber.org/zap"
	"net/http"
)

type IPhoneRegisterHandler struct {
	LoggerSugar *zap.SugaredLogger
}

func NewRegisterHandler(loggerSugar *zap.SugaredLogger) IPhoneRegisterHandler {
	return IPhoneRegisterHandler{
		LoggerSugar: loggerSugar,
	}
}

func (h *IPhoneRegisterHandler) RegisterScreen(w http.ResponseWriter, r *http.Request) {
	h.LoggerSugar.Warnw("health check")
}

func (h *IPhoneRegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	h.LoggerSugar.Warnw("health check")
}
