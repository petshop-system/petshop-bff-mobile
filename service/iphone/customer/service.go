package customer

import "github.com/petshop-system/petshop-bff-mobile/domain"

type IIphoneCustomerService interface {
	CustomerValidateCreate(ctxControl domain.ContextControl, customerCreate NewCustomerServiceDomain) error
	CustomerCreate(ctxControl domain.ContextControl, customerCreateService NewCustomerServiceDomain) (error, NewCustomerResponseServiceDomain)

	AddressValidateCreate(ctxControl domain.ContextControl, newAddress NewAddressServiceDomain) error
	AddressCreate(ctxControl domain.ContextControl, newAddress NewAddressServiceDomain) (error, NewAddressResponseServiceDomain)

	PhoneValidateCreate(ctxControl domain.ContextControl, newPhone NewPhoneServiceDomain) error
	PhoneCreate(ctxControl domain.ContextControl, newPhone NewPhoneServiceDomain) (error, NewPhoneResponseServiceDomain)
}
