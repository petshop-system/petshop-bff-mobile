package intergration

import (
	"bytes"
	"fmt"
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

func NewIntegration(client http.Client, loggerSugar *zap.SugaredLogger) Integration {
	return Integration{
		Client:      client,
		LoggerSugar: loggerSugar,
	}
}

func (integration *Integration) Post(url string, headers map[string]string, body *bytes.Buffer) (io.ReadCloser, error) {

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return integration.doRequest(req)
}

func (integration *Integration) Get(url string, headers map[string]string) (io.ReadCloser, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return integration.doRequest(req)
}

func (integration *Integration) doRequest(request *http.Request) (io.ReadCloser, error) {

	dump, _ := httputil.DumpRequest(request, true)
	integration.LoggerSugar.Infow("executing post request", "request", string(dump))
	resp, err := integration.Client.Do(request)
	if err != nil {
		return nil, err
	}

	dump, _ = httputil.DumpResponse(resp, true)
	integration.LoggerSugar.Infow("executing post request", "request", string(dump))

	if !utils.IsStatusCode2xx(resp.StatusCode) {
		return nil, fmt.Errorf("request failed. status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}
