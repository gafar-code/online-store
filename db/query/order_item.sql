-- name: CreateOrderItem :one

INSERT INTO
    order_items (
        product_id,
        qty,
        order_id
    )
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListOrderItemByOrderId :many

SELECT * FROM order_items WHERE order_id = $1 LIMIT $2 OFFSET $3;