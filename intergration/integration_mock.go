package intergration

import (
	"bytes"
	"github.com/petshop-system/petshop-bff-mobile/domain"
	"io"
)

type IntegrationMock struct {
	PostMock func(ctxControl domain.ContextControl, url string, headers map[string]string, body *bytes.Buffer) (io.ReadCloser, error)
	GetMock  func(ctxControl domain.ContextControl, url string, headers map[string]string) (io.ReadCloser, error)
}

func (im *IntegrationMock) Post(ctxControl domain.ContextControl, url string, headers map[string]string, body *bytes.Buffer) (io.ReadCloser, error) {
	if im.PostMock != nil {
		return im.PostMock(ctxControl, url, headers, body)
	}
	return nil, nil
}

func (im *IntegrationMock) Get(ctxControl domain.ContextControl, url string, headers map[string]string) (io.ReadCloser, error) {
	if im.GetMock != nil {
		return im.Get(ctxControl, url, headers)
	}
	return nil, nil
}
