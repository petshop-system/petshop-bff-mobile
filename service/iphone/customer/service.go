package customer

type IIphoneCustomerService interface {
	CustomerValidateCreate(customerCreate NewCustomerServiceDomain) error
	CustomerCreate(customerCreateService NewCustomerServiceDomain) (error, NewCustomerResponseServiceDomain)

	AddressValidateCreate(newAddress NewAddressServiceDomain) error
	AddressCreate(newAddress NewAddressServiceDomain) (error, NewAddressResponseServiceDomain)

	PhoneValidateCreate(newPhone NewPhoneServiceDomain) error
	PhoneCreate(newPhone NewPhoneServiceDomain) (error, NewPhoneResponseServiceDomain)
}
