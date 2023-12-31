-- name: CreateTransaction :one

INSERT INTO
    transactions (
        customer_id,
        status,
        issued_at,
        order_id
    )
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListTransactionByCustomerId :many

SELECT * FROM transactions WHERE customer_id = $3 LIMIT $1 OFFSET $2;