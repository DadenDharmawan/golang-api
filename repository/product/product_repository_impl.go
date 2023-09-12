package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DadenDharmawan/api-go/helper"
	"github.com/DadenDharmawan/api-go/model/entity"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repo *ProductRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product {

	query := "INSERT INTO products VALUES (null, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, product.Name, product.Category, product.Desc, product.Price, product.Qty, product.Image)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (repo *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product entity.Product) entity.Product {
	query := "UPDATE products SET name = ?, category = ?, description = ?, price = ?, qty = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, product.Name, product.Category, product.Desc, product.Price, product.Qty, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repo *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product entity.Product) {
	query := "DELETE FROM products WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, product.Id)
	helper.PanicIfError(err)
}

func (repo *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (entity.Product, error) {
	query := "SELECT id, name, category, description, price, qty, image FROM products WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	var product entity.Product
	if rows.Next() {
		// if data exist
		err := rows.Scan(&product.Id, &product.Name, &product.Category, &product.Desc, &product.Price, &product.Qty, &product.Image)
		helper.PanicIfError(err)
		return product, nil
	} else {
		// if data doesn't exist
		return product, errors.New("product is not found")
	}
}

func (repo *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Product {
	query := "SELECT id, name, category, description, price, qty, image FROM products"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		// iteration to get all data
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Category, &product.Desc, &product.Price, &product.Qty, &product.Image)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	return products
}

func (repo *ProductRepositoryImpl) FindbyCategory(ctx context.Context, tx *sql.Tx, category string) ([]entity.Product, error) {
	query := "SELECT id, name, category, description, price, qty, image FROM products WHERE category = ?"
	rows, err := tx.QueryContext(ctx, query, category)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Category, &product.Desc, &product.Price, &product.Qty, &product.Image)
		helper.PanicIfError(err)
		products = append(products, product)
	}

	if len(products) == 0 {
        return nil, errors.New("category is not found")
    }

	return products, nil
}
