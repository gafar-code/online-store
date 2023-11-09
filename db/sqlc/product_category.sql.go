// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: product_category.sql

package db

import (
	"context"
)

const createProductCategory = `-- name: CreateProductCategory :one

INSERT INTO
    product_categories (name)
VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateProductCategory(ctx context.Context, name string) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, createProductCategory, name)
	var i ProductCategory
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listProductCategory = `-- name: ListProductCategory :many

SELECT id, name FROM product_categories ORDER BY id LIMIT $1 OFFSET $2
`

type ListProductCategoryParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProductCategory(ctx context.Context, arg ListProductCategoryParams) ([]ProductCategory, error) {
	rows, err := q.db.QueryContext(ctx, listProductCategory, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductCategory{}
	for rows.Next() {
		var i ProductCategory
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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