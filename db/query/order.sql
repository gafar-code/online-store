-- name: CreateOrder :one

INSERT INTO
    orders (
        customer_id,
        amount,
        status,
        virtual_account_id,
        expired_at
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListOrderByCustomerId :many

SELECT * FROM orders WHERE customer_id = $1 LIMIT $2 OFFSET $3;

-- name: GetOrder :one

SELECT * FROM orders WHERE id = $1 LIMIT 1;

-- name: UpdateOrder :one

UPDATE orders SET status = $1 WHERE id = $2 RETURNING *;