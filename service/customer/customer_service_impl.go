package customer

import (
	"context"
	"database/sql"

	"github.com/DadenDharmawan/api-go/exception"
	"github.com/DadenDharmawan/api-go/helper"
	"github.com/DadenDharmawan/api-go/model/entity"
	customerWeb "github.com/DadenDharmawan/api-go/model/web"
	customerRepository "github.com/DadenDharmawan/api-go/repository/customer"
	"github.com/go-playground/validator/v10"
)

type customerServiceImpl struct {
	CustomerRepository customerRepository.CustomerRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCustonerService(customerRepository customerRepository.CustomerRepository, db *sql.DB, validate *validator.Validate) CustomerService {
	return &customerServiceImpl{
		CustomerRepository: customerRepository,
		DB: db,
		Validate: validate,
	}
}

func (service *customerServiceImpl) Insert(ctx context.Context, request customerWeb.CustomerInsertRequest) customerWeb.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := entity.Customer {
		Name: request.Name,
		Email: request.Email,
		Gender: request.Gender,
		Address: request.Address,
		Password: request.Password,
	}

	customer = service.CustomerRepository.Insert(ctx, tx, customer)
	return customerWeb.CustomerResponse(customer)
}

func (service *customerServiceImpl) Update(ctx context.Context, request customerWeb.CustomerUpdateRequest) customerWeb.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	customer.Name = request.Name
	customer.Email = request.Email
	customer.Gender = request.Gender
	customer.Address = request.Address
	customer.Password = request.Password
	customer = service.CustomerRepository.Update(ctx, tx, customer)

	return customerWeb.CustomerResponse(customer)
}

func (service *customerServiceImpl) Delete(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CustomerRepository.Delete(ctx, tx, customer)
}

func (service *customerServiceImpl) FindById(ctx context.Context, customerId int) customerWeb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return customerWeb.CustomerResponse(customer)
}

func (service *customerServiceImpl) FindAll(ctx context.Context) []customerWeb.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers := service.CustomerRepository.FindAll(ctx, tx)

	var customerResponses []customerWeb.CustomerResponse
	for _, customer := range customers {
		customerResponses = append(customerResponses, customerWeb.CustomerResponse(customer))
	}

	return customerResponses
}