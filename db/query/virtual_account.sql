-- name: CreateVirtualAccount :one

INSERT INTO
    virtual_accounts (
        name,
        description,
        rekening_number
    )
VALUES ($1, $2, $3)
RETURNING *;

-- name: ListVirtualAccount :many

SELECT * FROM virtual_accounts ORDER BY id;