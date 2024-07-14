package customer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/petshop-system/petshop-bff-mobile/domain"
	"github.com/petshop-system/petshop-bff-mobile/intergration"
	"github.com/petshop-system/petshop-bff-mobile/service/iphone/customer"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomerValidateCreate(t *testing.T) {

	loggerExample := zap.NewExample().Sugar()

	tests := []struct {
		Name           string
		Integration    intergration.IntegrationInterface
		RequestBody    NewCustomer
		ExpectedResult int
	}{
		{
			Name: "all fields are corrects. must return status code 200 (ok).",
			Integration: &intergration.IntegrationMock{
				PostMock: func(ctxControl domain.ContextControl, url string, headers map[string]string, body *bytes.Buffer) (io.ReadCloser, error) {
					return nil, nil
				},
			},
			RequestBody: NewCustomer{
				Name:       "Lechitz",
				ContractID: 1,
				PersonType: "individual",
				Document:   "612.178.283-89",
				Email:      "ferreirinha@petshop.com",
			},
			ExpectedResult: http.StatusOK,
		},
		{
			Name: "there is no document. must return status code 400 (bad request).",
			Integration: &intergration.IntegrationMock{
				PostMock: func(ctxControl domain.ContextControl, url string, headers map[string]string, body *bytes.Buffer) (io.ReadCloser, error) {
					return nil, fmt.Errorf("there are some errors")
				},
			},
			RequestBody: NewCustomer{
				Name:       "Lechitz",
				ContractID: 1,
				PersonType: "individual",
				Email:      "ferreirinha@petshop.com",
			},
			ExpectedResult: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

			host := "https://myhost.com"

			iphoneCustomerService := customer.NewIphoneCustomerService(loggerExample,
				tt.Integration, host)

			iPhoneCustomerHandler := NewIPhoneCustomerHandler(loggerExample, &iphoneCustomerService)

			body := new(bytes.Buffer)
			json.NewEncoder(body).Encode(tt.RequestBody)

			req := httptest.NewRequest(http.MethodPost, host, body)

			rc := req.Context()
			ctx := context.WithValue(rc, middleware.RequestIDHeader, "1234")

			w := httptest.NewRecorder()
			iPhoneCustomerHandler.CustomerValidateCreate(w, req.WithContext(ctx))

			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.ExpectedResult, res.StatusCode)

		})

	}

}
