// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: order_item.sql

package db

import (
	"context"
)

const createOrderItem = `-- name: CreateOrderItem :one

INSERT INTO
    order_items (
        category_id,
        name,
        image_url,
        description,
        price,
        qty,
        product_id,
        order_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, category_id, name, image_url, description, price, qty, product_id, order_id, created_at
`

type CreateOrderItemParams struct {
	CategoryID  int64  `json:"category_id"`
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Qty         int64  `json:"qty"`
	ProductID   int64  `json:"product_id"`
	OrderID     int64  `json:"order_id"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.CategoryID,
		arg.Name,
		arg.ImageUrl,
		arg.Description,
		arg.Price,
		arg.Qty,
		arg.ProductID,
		arg.OrderID,
	)
	var i OrderItem
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.ImageUrl,
		&i.Description,
		&i.Price,
		&i.Qty,
		&i.ProductID,
		&i.OrderID,
		&i.CreatedAt,
	)
	return i, err
}

const listOrderItemByOrderId = `-- name: ListOrderItemByOrderId :many

SELECT id, category_id, name, image_url, description, price, qty, product_id, order_id, created_at FROM order_items WHERE order_id = $1
`

func (q *Queries) ListOrderItemByOrderId(ctx context.Context, orderID int64) ([]OrderItem, error) {
	rows, err := q.db.QueryContext(ctx, listOrderItemByOrderId, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []OrderItem{}
	for rows.Next() {
		var i OrderItem
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.ImageUrl,
			&i.Description,
			&i.Price,
			&i.Qty,
			&i.ProductID,
			&i.OrderID,
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
