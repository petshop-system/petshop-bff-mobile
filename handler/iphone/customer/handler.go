package customer

import (
	"github.com/petshop-system/petshop-bff-mobile/service/iphone/customer"
	"go.uber.org/zap"
)

const (
	InvalidationCreateCustomer = "there are some invalid customer info"
	SuccessToCreateCustomer    = "success to create the customer"
	InvalidationCreateAddress  = "there are some invalid address info"
	InvalidationCreatePhone    = "there are some invalid phone info"
)

type IPhoneCustomerHandler struct {
	loggerSugar   *zap.SugaredLogger
	iphoneService customer.IIphoneCustomerService
}

func NewIPhoneCustomerHandler(loggerSugar *zap.SugaredLogger, iphoneService customer.IIphoneCustomerService) IPhoneCustomerHandler {
	return IPhoneCustomerHandler{
		loggerSugar:   loggerSugar,
		iphoneService: iphoneService,
	}
}
