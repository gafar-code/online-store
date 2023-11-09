-- name: CreateOrder :one

INSERT INTO
    orders (
        customer_id,
        virtual_account_id
    )
VALUES ($1, $2)
RETURNING *;

-- name: ListOrderByCustomerId :many

SELECT * FROM orders WHERE customer_id = $1 LIMIT $2 OFFSET $3;