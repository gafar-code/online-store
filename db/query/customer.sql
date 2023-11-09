-- name: CreateCustomer :one

INSERT INTO
    customers (
        name,
        email,
        password,
        address,
        token
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetCustomer :one

SELECT * FROM customers WHERE id = $1 LIMIT 1;