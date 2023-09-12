package customer

import (
	"context"
	"database/sql"
	"errors"

	"github.com/DadenDharmawan/api-go/helper"
	"github.com/DadenDharmawan/api-go/model/entity"
)

type customerRepositoryImpl struct {
}

func NewcustomerRepository() CustomerRepository {
	return &customerRepositoryImpl{}
}

func (repo *customerRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer {
	query := "INSERT INTO customers VALUES (null, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, customer.Name, customer.Email, customer.Gender, customer.Address, customer.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	customer.Id = int(id)
	return customer
}

func (repo *customerRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, customer entity.Customer) entity.Customer {
	query := "UPDATE customers SET name = ?, email = ?, gender = ?, address = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, customer.Name, customer.Email, customer.Gender, customer.Address, customer.Id)
	helper.PanicIfError(err)

	return customer
}

func (repo *customerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customer entity.Customer) {
	query := "DELETE FROM customers WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, customer.Id)
	helper.PanicIfError(err)
}

func (repo *customerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customerId int) (entity.Customer, error) {
	query := "SELECT id, name, email, gender, address, password FROM customers WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	var customer entity.Customer
	if rows.Next() {
		// if data exist
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Gender, &customer.Address, &customer.Password)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		// if data doesn't exist
		return customer, errors.New("customer is not found")
	}
}

func (repo *customerRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Customer {
	query := "SELECT id, name, email, gender, address, password FROM customers"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		// iteration to get all data
		var customer entity.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Gender, &customer.Address, &customer.Password)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}

	return customers
}