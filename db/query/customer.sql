-- name: CreateCustomer :one

INSERT INTO
    customers (
        name,
        email,
        password,
        address
    )
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetCustomerByEmail :one

SELECT * FROM customers WHERE email = $1 LIMIT 1;