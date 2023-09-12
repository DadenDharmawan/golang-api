package customer

import (
	"context"

	"github.com/DadenDharmawan/api-go/model/web"
)

type CustomerService interface {
	Insert(ctx context.Context, request web.CustomerInsertRequest) web.CustomerResponse
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse
	Delete(ctx context.Context, customerId int)
	FindById(ctx context.Context, customerId int) web.CustomerResponse
	FindAll(ctx context.Context) []web.CustomerResponse
}