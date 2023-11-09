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

-- name: GetProductDetail :one

SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: GetProductByCategoryId :many

SELECT * FROM products WHERE category_id = $1 LIMIT $2 OFFSET $3;

-- name: ListProduct :many

SELECT * FROM products ORDER BY id LIMIT $1 OFFSET $2;