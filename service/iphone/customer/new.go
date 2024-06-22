package customer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/petshop-system/petshop-bff-mobile/intergration"
	"go.uber.org/zap"
	"io"
)

type IphoneService struct {
	loggerSugar         *zap.SugaredLogger
	CustomerIntegration *intergration.Integration
	APIGatewayHost      string
}

const (
	CustomerValidateCreateURN = "/customer/validate-create"
	CustomerCreateURN         = "/customer/create"
	AddressValidateCreateURN  = "/address/validate-create"
	AddressCreateURN          = "/address/create"
	PhoneValidateCreateURN    = "/phone/validate-create"
	PhoneCreateURN            = "/phone/create"
)

func NewIphoneCustomerService(loggerSugar *zap.SugaredLogger, customerIntegration *intergration.Integration,
	apiGatewayHost string) IphoneService {

	return IphoneService{
		loggerSugar:         loggerSugar,
		CustomerIntegration: customerIntegration,
		APIGatewayHost:      apiGatewayHost,
	}
}

type NewCustomerServiceDomain struct {
	Name       string
	Email      string
	Document   string
	PersonType string
	ContractID int64
	AddressID  int64
}

type newCustomerRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Document   string `json:"document"`
	PersonType string `json:"person_type"`
	ContractID int64  `json:"contract_id"`
	AddressID  int64  `json:"address_id"`
}

type NewCustomerResponseServiceDomain struct {
	ID         int64
	Name       string
	Email      string
	Document   string
	PersonType string
	ContractID int64
	AddressID  int64
}

type NewAddressServiceDomain struct {
	Street  string
	Number  string
	Zipcode string
}

type newAddressRequest struct {
	Street  string `json:"street"`
	Number  string `json:"number"`
	Zipcode string `json:"zipcode"`
}

type NewAddressResponseServiceDomain struct {
	ID      int64
	Street  string
	Number  string
	Zipcode string
}

type NewPhoneServiceDomain struct {
	Number    string
	CodeArea  string
	PhoneType string // mobile, particular, work
	UserID    int64
	UserType  string // contract, employee, customer
}

type newPhoneRequest struct {
	Number    string `json:"number"`
	CodeArea  string `json:"code_area"`
	PhoneType string `json:"phone_type"` // mobile, particular, work
	UserID    int64  `json:"user_id"`
	UserType  string `json:"user_type"` // contract, employee, customer
}

type NewPhoneResponseServiceDomain struct {
	ID        int64
	Number    string
	CodeArea  string
	PhoneType string
	UserID    int64
	UserType  string // contract, employee, customer
}

func (service *IphoneService) getURI(path string) string {
	return fmt.Sprintf("%s%s", service.APIGatewayHost, path)
}

func (service *IphoneService) CustomerValidateCreate(customerCreate NewCustomerServiceDomain) error {

	var newCustomerReq newCustomerRequest
	copier.Copy(&newCustomerReq, &customerCreate)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(newCustomerReq); err != nil {
		return err
	}

	url := service.getURI(CustomerValidateCreateURN)
	_, err := service.CustomerIntegration.Post(url, nil, &body)

	return err
}

func (service *IphoneService) CustomerCreate(customerCreateService NewCustomerServiceDomain) (error, NewCustomerResponseServiceDomain) {

	var newCustomerReq newCustomerRequest
	copier.Copy(&newCustomerReq, &customerCreateService)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(newCustomerReq); err != nil {
		return err, NewCustomerResponseServiceDomain{}
	}

	url := service.getURI(CustomerCreateURN)
	post, err := service.CustomerIntegration.Post(url, nil, &body)
	if err != nil {
		return err, NewCustomerResponseServiceDomain{}
	}

	var newServiceResponse NewCustomerResponseServiceDomain
	bytes, err := io.ReadAll(post)
	json.Unmarshal(bytes, &newServiceResponse)

	return nil, newServiceResponse
}

func (service *IphoneService) AddressValidateCreate(newAddress NewAddressServiceDomain) error {

	var newAddressReq newAddressRequest
	copier.Copy(&newAddressReq, newAddress)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(newAddressReq); err != nil {
		return err
	}

	url := service.getURI(AddressValidateCreateURN)
	_, err := service.CustomerIntegration.Post(url, nil, &body)

	return err
}

func (service *IphoneService) AddressCreate(newAddress NewAddressServiceDomain) (error, NewAddressResponseServiceDomain) {

	var newAddressReq newAddressRequest
	copier.Copy(&newAddressReq, newAddress)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(newAddressReq); err != nil {
		return err, NewAddressResponseServiceDomain{}
	}

	url := service.getURI(AddressCreateURN)
	post, err := service.CustomerIntegration.Post(url, nil, &body)
	if err != nil {
		return err, NewAddressResponseServiceDomain{}
	}

	var newServiceResponse NewAddressResponseServiceDomain
	bytes, err := io.ReadAll(post)
	json.Unmarshal(bytes, &newServiceResponse)

	return nil, newServiceResponse
}

func (service *IphoneService) PhoneValidateCreate(newPhone NewPhoneServiceDomain) error {

	var newPhoneReq newPhoneRequest
	copier.Copy(&newPhoneReq, &newPhone)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(newPhoneReq); err != nil {
		return err
	}

	url := service.getURI(PhoneValidateCreateURN)
	_, err := service.CustomerIntegration.Post(url, nil, &body)

	return err
}

func (service *IphoneService) PhoneCreate(newPhone NewPhoneServiceDomain) (error, NewPhoneResponseServiceDomain) {

	var newPhoneReq newPhoneRequest
	copier.Copy(&newPhoneReq, &newPhone)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(newPhoneReq); err != nil {
		return err, NewPhoneResponseServiceDomain{}
	}

	url := service.getURI(PhoneCreateURN)
	post, err := service.CustomerIntegration.Post(url, nil, &body)
	if err != nil {
		return err, NewPhoneResponseServiceDomain{}
	}

	var newServiceResponse NewPhoneResponseServiceDomain
	bytes, err := io.ReadAll(post)
	json.Unmarshal(bytes, &newServiceResponse)

	return nil, newServiceResponse
}
