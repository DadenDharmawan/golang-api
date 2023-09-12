package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DadenDharmawan/api-go/helper"
	"github.com/DadenDharmawan/api-go/model/web"
	customerService "github.com/DadenDharmawan/api-go/service/customer"
	"github.com/julienschmidt/httprouter"
)

type customerControllerImpl struct {
	CustomerService customerService.CustomerService
}

func NewCustomerController(cutomerService customerService.CustomerService) CustomerController {
	return &customerControllerImpl{
		CustomerService: cutomerService,
	}
}

func (controller *customerControllerImpl) Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var customerInsertRequest web.CustomerInsertRequest

	err := decoder.Decode(&customerInsertRequest)
	helper.PanicIfError(err)

	customerResponse := controller.CustomerService.Insert(r.Context(), customerInsertRequest)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *customerControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var customerUpdateRequest web.CustomerUpdateRequest

	err := decoder.Decode(&customerUpdateRequest)
	helper.PanicIfError(err)

	customerId := p.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerUpdateRequest.Id = id

	customerResponse := controller.CustomerService.Update(r.Context(), customerUpdateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *customerControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	customerId := p.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	controller.CustomerService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: nil,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *customerControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	customerId := p.ByName("customerId")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerResponse := controller.CustomerService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *customerControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	customerResponse := controller.CustomerService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Status: "OK",
		Data: customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}