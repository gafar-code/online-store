-- name: CreateOrderItem :one

INSERT INTO
    order_items (
        product_id,
        product_price,
        qty,
        order_id
    )
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListOrderItemByOrderId :many

SELECT * FROM order_items WHERE order_id = $1 LIMIT $2 OFFSET $3;