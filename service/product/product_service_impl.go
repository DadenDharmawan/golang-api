package product

import (
	"context"
	"database/sql"
	"io"
	"os"
	"strings"

	"github.com/DadenDharmawan/api-go/exception"
	"github.com/DadenDharmawan/api-go/helper"
	"github.com/DadenDharmawan/api-go/model/entity"
	productWeb "github.com/DadenDharmawan/api-go/model/web"
	productRepository "github.com/DadenDharmawan/api-go/repository/product"
	"github.com/go-playground/validator/v10"
)

type productServiceImpl struct {
	ProductRepository productRepository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository productRepository.ProductRepository, db *sql.DB, validate *validator.Validate) ProductService {
	return &productServiceImpl{
		ProductRepository: productRepository,
		DB:                db,
		Validate:          validate,
	}
}

func (service *productServiceImpl) Insert(ctx context.Context, request productWeb.ProductInsertRequest) productWeb.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var responses []entity.Product
	for _, image := range request.Image {
		file, _ := image.Open()

		tempFile, err := os.CreateTemp("public", "image-*.jpg")
		helper.PanicIfError(err)
		defer tempFile.Close()

		fileBytes, err := io.ReadAll(file)
		helper.PanicIfError(err)

		tempFile.Write(fileBytes)

		fileName := tempFile.Name()
		newFileName := strings.Split(fileName, "\\")

		product := entity.Product{
			Name:     request.Name,
			Category: request.Category,
			Desc:     request.Desc,
			Price:    request.Price,
			Qty:      request.Qty,
			Image:    newFileName[1],
		}

		product = service.ProductRepository.Insert(ctx, tx, product)
		responses = append(responses, product)
	}

	return productWeb.ProductResponse(responses[0])
}

func (service *productServiceImpl) Update(ctx context.Context, request productWeb.ProductUpdateRequest) productWeb.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	product.Name = request.Name
	product.Category = request.Category
	product.Desc = request.Desc
	product.Price = request.Price
	product.Qty = request.Qty
	product = service.ProductRepository.Update(ctx, tx, product)

	return productWeb.ProductResponse(product)
}

func (service *productServiceImpl) Delete(ctx context.Context, productId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, product)
}

func (service *productServiceImpl) FindById(ctx context.Context, productId int) productWeb.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return productWeb.ProductResponse(product)
}

func (service *productServiceImpl) FindAll(ctx context.Context) []productWeb.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAll(ctx, tx)

	var productResponses []productWeb.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, productWeb.ProductResponse(product))
	}

	return productResponses
}

func (service *productServiceImpl) FindbyCategory(ctx context.Context, category string) []productWeb.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories, err := service.ProductRepository.FindbyCategory(ctx, tx, category)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	var productsResponses []productWeb.ProductResponse
	for _, cattegory := range categories {
		productsResponses = append(productsResponses, productWeb.ProductResponse(cattegory))
	}

	return productsResponses
}
