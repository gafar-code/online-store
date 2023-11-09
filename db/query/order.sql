-- name: CreateOrder :one

INSERT INTO
    orders (
        customer_id,
        status,
        virtual_account_id
    )
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListOrderByCustomerId :many

SELECT * FROM orders WHERE customer_id = $1 LIMIT $2 OFFSET $3;