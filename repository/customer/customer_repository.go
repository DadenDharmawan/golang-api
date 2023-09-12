package customer

import (
	"context"
	"database/sql"

	"github.com/DadenDharmawan/api-go/model/entity"
)

type CustomerRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer
	Update(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer
	Delete(ctx context.Context, tx *sql.Tx, customer entity.Customer)
	FindById(ctx context.Context, tx *sql.Tx, customerId int) (entity.Customer, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Customer
}