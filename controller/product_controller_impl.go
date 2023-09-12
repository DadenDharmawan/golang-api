package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DadenDharmawan/api-go/helper"
	"github.com/DadenDharmawan/api-go/model/web"
	productService "github.com/DadenDharmawan/api-go/service/product"
	"github.com/julienschmidt/httprouter"
)

type productControllerImpl struct {
	ProductController productService.ProductService
}

func NewProductController(ProductService productService.ProductService) ProductController {
	return &productControllerImpl{
		ProductController: ProductService,
	}
}

func (controller *productControllerImpl) Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseMultipartForm(10 * 1024 * 1024)
	// decoder := json.NewDecoder(r.Body)
	var productInsertRequest web.ProductInsertRequest
	// err := decoder.Decode(&productInsertRequest)
	// helper.PanicIfError(err)

	productInsertRequest.Name = r.PostFormValue("name")
	productInsertRequest.Category = r.PostFormValue("category")

	if r.PostFormValue("desc") != "" {
		productInsertRequest.Desc = r.PostFormValue("desc")
	} else {
		productInsertRequest.Desc = ""
	}

	price, err := strconv.Atoi(r.PostFormValue("price"))
	helper.PanicIfError(err)
	productInsertRequest.Price = price

	qty, err := strconv.Atoi(r.PostFormValue("qty"))
	helper.PanicIfError(err)

	productInsertRequest.Qty = qty

	if r.MultipartForm.File["image"] != nil {
		productInsertRequest.Image = r.MultipartForm.File["image"]
	} else {
		productInsertRequest.Image = nil
	}

	customerResponse := controller.ProductController.Insert(r.Context(), productInsertRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   customerResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

// func (controller *productControllerImpl) Insert(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	contentType := r.Header.Get("Content-Type")

// 	var productInsertRequest web.ProductInsertRequest

// 	switch contentType {

// 	case "multipart":
// 		err := r.ParseMultipartForm(10 * 1024 * 1024) // Misalnya, batasan ukuran 10 MB
// 		helper.PanicIfError(err)

// 		productInsertRequest.Name = r.PostFormValue("name")
// 		productInsertRequest.Category = r.PostFormValue("category")
// 		productInsertRequest.Desc = r.PostFormValue("desc")
// 		price, err := strconv.Atoi(r.PostFormValue("price"))
// 		helper.PanicIfError(err)
// 		productInsertRequest.Price = price
// 		qty, err := strconv.Atoi(r.PostFormValue("qty"))
// 		helper.PanicIfError(err)
// 		productInsertRequest.Qty = qty
// 		productInsertRequest.Image = r.MultipartForm.File["image"]
// 	case "application/json":
// 		decoder := json.NewDecoder(r.Body)
// 		err := decoder.Decode(&productInsertRequest)
// 		helper.PanicIfError(err)
// 	}

// 	customerResponse := controller.ProductController.Insert(r.Context(), productInsertRequest)
// 	webResponse := web.WebResponse{
// 		Code:   http.StatusOK,
// 		Status: "OK",
// 		Data:   customerResponse,
// 	}

// 	helper.WriteToResponseBody(w, webResponse)
// }

func (controller *productControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var productUpdateRequest web.ProductUpdateRequest

	err := decoder.Decode(&productUpdateRequest)
	helper.PanicIfError(err)

	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	ProductResponse := controller.ProductController.Update(r.Context(), productUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *productControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	controller.ProductController.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *productControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productId := p.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := controller.ProductController.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *productControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productResponse := controller.ProductController.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *productControllerImpl) FindbyCategory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productCategory := p.ByName("productCategory")

	productResponse := controller.ProductController.FindbyCategory(r.Context(), productCategory)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
