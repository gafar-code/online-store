-- name: CreateOrderProof :one

INSERT INTO
    order_proof (
        order_id,
        name_holder,
        rekening_number,
        image_url
    )
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOrderProof :one

SELECT * FROM order_proof WHERE order_id = $1 LIMIT $1;