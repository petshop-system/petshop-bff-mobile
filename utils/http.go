package utils

import (
	"net/http"
	"time"
)

func NewHttpClient(timeout time.Duration, maxIdleConns int, idleConnTimeout time.Duration) http.Client {
	return http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			MaxIdleConns:       maxIdleConns,
			IdleConnTimeout:    idleConnTimeout,
			DisableCompression: true,
		},
	}
}

func IsStatusCode2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
