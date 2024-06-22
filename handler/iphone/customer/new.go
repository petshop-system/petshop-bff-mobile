package customer

import (
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/petshop-system/petshop-bff-mobile/service/iphone/customer"
	"github.com/petshop-system/petshop-bff-mobile/utils"
	"net/http"
)

type NewCustomerCreateRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Document   string `json:"document"`
	PersonType string `json:"person_type"`
	ContractID int64  `json:"contract_id"`
	Address    struct {
		Street  string `json:"street"`
		Number  string `json:"number"`
		Zipcode string `json:"zipcode"`
	} `json:"address"`
	Phone struct {
		Number    string `json:"number"`
		CodeArea  string `json:"code_area"`
		PhoneType string `json:"phone_type"`
	} `json:"phone"`
}

type NewCustomerCreateResponse struct {
	ID         int64
	Name       string
	Email      string
	Document   string
	PersonType string
	ContractID int64
	Address    struct {
		ID      int64
		Street  string
		Number  string
		Zipcode string
	}
	Phone struct {
		ID        int64
		Number    string
		CodeArea  string
		PhoneType string
	}
}

func (h *IPhoneCustomerHandler) CreateScreen(w http.ResponseWriter, r *http.Request) {

	h.loggerSugar.Warnw("health check")
}

func (h *IPhoneCustomerHandler) Create(w http.ResponseWriter, r *http.Request) {

	var customerCreateRequest NewCustomerCreateRequest
	json.NewDecoder(r.Body).Decode(&customerCreateRequest)

	// validate address
	var newAddressServiceDomain customer.NewAddressServiceDomain
	_ = copier.Copy(&newAddressServiceDomain, &customerCreateRequest.Address)
	if err := h.iphoneService.AddressValidateCreate(newAddressServiceDomain); err != nil {
		h.loggerSugar.Errorw(InvalidationCreateAddress, "error", err.Error())
		response := utils.ObjectResponse(err.Error(), InvalidationCreateAddress)
		utils.ResponseReturn(w, http.StatusBadRequest, response.Bytes())
		return
	}

	// validate customer
	var newCustomerServiceDomain customer.NewCustomerServiceDomain
	_ = copier.Copy(&newCustomerServiceDomain, &customerCreateRequest)
	if err := h.iphoneService.CustomerValidateCreate(newCustomerServiceDomain); err != nil {
		h.loggerSugar.Errorw(InvalidationCreateCustomer, "error", err.Error())
		response := utils.ObjectResponse(err.Error(), InvalidationCreateCustomer)
		utils.ResponseReturn(w, http.StatusBadRequest, response.Bytes())
		return
	}

	// validate phone
	var newPhoneServiceDomain customer.NewPhoneServiceDomain
	_ = copier.Copy(&newPhoneServiceDomain, &customerCreateRequest.Phone)
	newPhoneServiceDomain.UserType = "customer"
	if err := h.iphoneService.PhoneValidateCreate(newPhoneServiceDomain); err != nil {
		h.loggerSugar.Errorw(InvalidationCreatePhone, "error", err.Error())
		response := utils.ObjectResponse(err.Error(), InvalidationCreatePhone)
		utils.ResponseReturn(w, http.StatusBadRequest, response.Bytes())
		return
	}

	// create address
	err, newAddressResponseServiceDomain := h.iphoneService.AddressCreate(newAddressServiceDomain)
	if err != nil {
		h.loggerSugar.Errorw(InvalidationCreateAddress, "error", err.Error())
		response := utils.ObjectResponse(err.Error(), InvalidationCreateAddress)
		utils.ResponseReturn(w, http.StatusInternalServerError, response.Bytes())
	}

	// create customer
	newCustomerServiceDomain.AddressID = newAddressResponseServiceDomain.ID
	err, newCustomerResponseServiceDomain := h.iphoneService.CustomerCreate(newCustomerServiceDomain)
	if err != nil {
		h.loggerSugar.Errorw(InvalidationCreateCustomer, "error", err.Error())
		response := utils.ObjectResponse(err.Error(), InvalidationCreateCustomer)
		utils.ResponseReturn(w, http.StatusInternalServerError, response.Bytes())
		return
	}

	// create phone
	newPhoneServiceDomain.UserID = newCustomerResponseServiceDomain.ID
	err, newPhoneResponseServiceDomain := h.iphoneService.PhoneCreate(newPhoneServiceDomain)
	if err != nil {
		h.loggerSugar.Errorw(InvalidationCreatePhone, "error", err.Error())
		response := utils.ObjectResponse(err.Error(), InvalidationCreatePhone)
		utils.ResponseReturn(w, http.StatusInternalServerError, response.Bytes())
		return
	}

	var newCustomerCreateResponse NewCustomerCreateResponse
	_ = copier.Copy(&newCustomerCreateResponse, &newCustomerResponseServiceDomain)
	_ = copier.Copy(&newCustomerCreateResponse.Address, &newAddressResponseServiceDomain)
	_ = copier.Copy(&newCustomerCreateResponse.Phone, &newPhoneResponseServiceDomain)

	h.loggerSugar.Infow(SuccessToCreateCustomer, "id", newCustomerCreateResponse.ID,
		"document", newCustomerCreateResponse.Document, "person_type", newCustomerCreateResponse.PersonType)

	response := utils.ObjectResponse(newCustomerCreateResponse, SuccessToCreateCustomer)
	utils.ResponseReturn(w, http.StatusCreated, response.Bytes())

}
