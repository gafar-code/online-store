// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: cart.sql

package db

import (
	"context"
)

const createCart = `-- name: CreateCart :one

INSERT INTO
    carts (customer_id, product_id, qty)
VALUES ($1, $2, $3)
RETURNING id, customer_id, product_id, qty, created_at
`

type CreateCartParams struct {
	CustomerID int64 `json:"customer_id"`
	ProductID  int64 `json:"product_id"`
	Qty        int64 `json:"qty"`
}

func (q *Queries) CreateCart(ctx context.Context, arg CreateCartParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, createCart, arg.CustomerID, arg.ProductID, arg.Qty)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.ProductID,
		&i.Qty,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCart = `-- name: DeleteCart :exec

DELETE FROM carts WHERE id = $1
`

func (q *Queries) DeleteCart(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCart, id)
	return err
}

const getCartByCustomerId = `-- name: GetCartByCustomerId :many

SELECT id, customer_id, product_id, qty, created_at FROM carts WHERE customer_id = $1 LIMIT 1
`

func (q *Queries) GetCartByCustomerId(ctx context.Context, customerID int64) ([]Cart, error) {
	rows, err := q.db.QueryContext(ctx, getCartByCustomerId, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Cart{}
	for rows.Next() {
		var i Cart
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.ProductID,
			&i.Qty,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCart = `-- name: UpdateCart :one

UPDATE carts SET qty = $2 WHERE id = $1 RETURNING id, customer_id, product_id, qty, created_at
`

type UpdateCartParams struct {
	ID  int64 `json:"id"`
	Qty int64 `json:"qty"`
}

func (q *Queries) UpdateCart(ctx context.Context, arg UpdateCartParams) (Cart, error) {
	row := q.db.QueryRowContext(ctx, updateCart, arg.ID, arg.Qty)
	var i Cart
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.ProductID,
		&i.Qty,
		&i.CreatedAt,
	)
	return i, err
}
