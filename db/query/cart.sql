-- name: CreateCart :one

INSERT INTO
    carts (customer_id, product_id, qty)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCartByCustomerId :many

SELECT * FROM carts WHERE customer_id = $1 LIMIT $2 OFFSET $3;

-- name: GetCart :one

SELECT * FROM carts WHERE id = $1 LIMIT 1;

-- name: UpdateCart :one

UPDATE carts SET qty = $2 WHERE id = $1 RETURNING *;

-- name: DeleteCart :exec

DELETE FROM carts WHERE id = $1;