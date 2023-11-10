-- name: CreateOrderItem :one

INSERT INTO
    order_items (
        category_id,
        name,
        image_url,
        description,
        price,
        qty,
        product_id,
        order_id
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: ListOrderItemByOrderId :many

SELECT * FROM order_items WHERE order_id = $1;