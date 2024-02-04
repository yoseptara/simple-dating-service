-- name: CreateUser :one
INSERT INTO
    users (username, password, fullname, email) VALUES (
        $1, $2, $3, $4
    )
    RETURNING *;