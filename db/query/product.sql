-- name: CreateProduct :one

INSERT INTO
    products (
        category_id,
        name,
        image_url,
        description,
        qty,
        price
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProduct :one

SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: GetProductByCategoryId :many

SELECT * FROM products WHERE category_id = $1 LIMIT $2 OFFSET $3;