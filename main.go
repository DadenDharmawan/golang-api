package main

import (
	"net/http"

	"github.com/DadenDharmawan/api-go/apps"
	"github.com/DadenDharmawan/api-go/controller"
	"github.com/DadenDharmawan/api-go/exception"
	"github.com/DadenDharmawan/api-go/helper"
	"github.com/DadenDharmawan/api-go/middleware"
	customerRepository "github.com/DadenDharmawan/api-go/repository/customer"
	productRepository "github.com/DadenDharmawan/api-go/repository/product"
	customerService "github.com/DadenDharmawan/api-go/service/customer"
	productService "github.com/DadenDharmawan/api-go/service/product"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := apps.ConnectionToDatabase()
	validate := validator.New()

	customerRepository := customerRepository.NewcustomerRepository()
	customerService := customerService.NewCustonerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	productRepository := productRepository.NewProductRepository()
	productService := productService.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := httprouter.New()

	// end point customers
	router.POST("/api/customers", customerController.Insert)
	router.PUT("/api/customers/:customerId", customerController.Update)
	router.DELETE("/api/customers/:customerId", customerController.Delete)
	router.GET("/api/customers/:customerId", customerController.FindById)
	router.GET("/api/customers", customerController.FindAll)

	// end point products
	router.POST("/api/products", productController.Insert)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)
	router.GET("/api/products/:productId", productController.FindById)
	router.GET("/api/products", productController.FindAll)
	router.GET("/api/categories/:productCategory", productController.FindbyCategory)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}