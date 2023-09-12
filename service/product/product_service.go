package product

import (
	"context"

	"github.com/DadenDharmawan/api-go/model/web"
)

type ProductService interface {
	Insert(ctx context.Context, request web.ProductInsertRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, productId int)
	FindById(ctx context.Context, productId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
	FindbyCategory(ctx context.Context, productCategory string) []web.ProductResponse
}