-- name: CreateProductCategory :one

INSERT INTO
    product_categories (name)
VALUES ($1)
RETURNING *;

-- name: ListProductCategory :many

SELECT * FROM product_categories ORDER BY id LIMIT $1 OFFSET $2;