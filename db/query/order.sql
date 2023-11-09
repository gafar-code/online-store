-- name: CreateOrder :one

INSERT INTO
    orders (
        customer_id,
        virtual_account_id
    )
VALUES ($1, $2)
RETURNING *;

-- name: ListOrderByCustomerId :many

SELECT * FROM orders ORDER BY customer_id LIMIT $1 OFFSET $2;