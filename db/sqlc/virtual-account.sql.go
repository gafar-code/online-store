// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: virtual-account.sql

package db

import (
	"context"
)

const createVirtualAccount = `-- name: CreateVirtualAccount :one

INSERT INTO
    virtual_accounts (
        name,
        description,
        rekening_number
    )
VALUES ($1, $2, $3)
RETURNING id, name, description, rekening_number, created_at
`

type CreateVirtualAccountParams struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	RekeningNumber int64  `json:"rekening_number"`
}

func (q *Queries) CreateVirtualAccount(ctx context.Context, arg CreateVirtualAccountParams) (VirtualAccount, error) {
	row := q.db.QueryRowContext(ctx, createVirtualAccount, arg.Name, arg.Description, arg.RekeningNumber)
	var i VirtualAccount
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.RekeningNumber,
		&i.CreatedAt,
	)
	return i, err
}

const listVirtualAccount = `-- name: ListVirtualAccount :many

SELECT id, name, description, rekening_number, created_at FROM virtual_accounts ORDER BY id LIMIT $1 OFFSET $2
`

type ListVirtualAccountParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListVirtualAccount(ctx context.Context, arg ListVirtualAccountParams) ([]VirtualAccount, error) {
	rows, err := q.db.QueryContext(ctx, listVirtualAccount, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []VirtualAccount{}
	for rows.Next() {
		var i VirtualAccount
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.RekeningNumber,
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
