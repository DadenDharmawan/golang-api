package product

import (
	"context"
	"database/sql"

	"github.com/DadenDharmawan/api-go/model/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product
	Update(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product
	Delete(ctx context.Context, tx *sql.Tx, product entity.Product)
	FindById(ctx context.Context, tx *sql.Tx, productId int) (entity.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Product
	FindbyCategory(ctx context.Context, tx *sql.Tx, category string) ([]entity.Product, error)
}