package utils

import (
	"context"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/petshop-system/petshop-bff-mobile/domain"
	"net/http"
)

func GetRequestIDAndContext(r *http.Request) (string, domain.ContextControl) {

	requestID := r.Context().Value(middleware.RequestIDHeader)

	ctxControl := domain.ContextControl{
		Context: context.WithValue(context.Background(),
			middleware.RequestIDHeader, requestID),
	}

	return requestID.(string), ctxControl
}
