// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: customer.sql

package db

import (
	"context"
)

const createCustomer = `-- name: CreateCustomer :one

INSERT INTO
    customers (
        name,
        email,
        password,
        address
    )
VALUES ($1, $2, $3, $4)
RETURNING id, name, email, password, address, created_at
`

type CreateCustomerParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.Address,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const getCustomerByEmail = `-- name: GetCustomerByEmail :one

SELECT id, name, email, password, address, created_at FROM customers WHERE email = $1 LIMIT 1
`

func (q *Queries) GetCustomerByEmail(ctx context.Context, email string) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByEmail, email)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}
