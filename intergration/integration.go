package intergration

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/petshop-system/petshop-bff-mobile/domain"
	"github.com/petshop-system/petshop-bff-mobile/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/http/httputil"
)

type Integration struct {
	Client      http.Client
	LoggerSugar *zap.SugaredLogger
}

type IntegrationInterface interface {
	Post(ctxControl domain.ContextControl, url string, headers map[string]string, body *bytes.Buffer) (io.ReadCloser, error)
	Get(ctxControl domain.ContextControl, url string, headers map[string]string) (io.ReadCloser, error)
}

func NewIntegration(client http.Client, loggerSugar *zap.SugaredLogger) Integration {
	return Integration{
		Client:      client,
		LoggerSugar: loggerSugar,
	}
}

func (integration *Integration) Post(ctxControl domain.ContextControl, url string, headers map[string]string, body *bytes.Buffer) (io.ReadCloser, error) {

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return integration.doRequest(ctxControl, req)
}

func (integration *Integration) Get(ctxControl domain.ContextControl, url string, headers map[string]string) (io.ReadCloser, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return integration.doRequest(ctxControl, req)
}

func (integration *Integration) doRequest(ctxControl domain.ContextControl, request *http.Request) (io.ReadCloser, error) {

	requestID := ctxControl.Context.Value(middleware.RequestIDKey)
	logger := integration.LoggerSugar.With("request_id", requestID)

	dump, _ := httputil.DumpRequest(request, true)
	logger.Infow("executing post request", "request", string(dump))

	resp, err := integration.Client.Do(request)
	if err != nil {
		return nil, err
	}

	dump, _ = httputil.DumpResponse(resp, true)
	logger.Infow("executing post request", "request", string(dump))

	if !utils.IsStatusCode2xx(resp.StatusCode) {
		return nil, fmt.Errorf("request failed. status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}
