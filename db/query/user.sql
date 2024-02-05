-- name: CreateUser :one
INSERT INTO
    users (username, password, fullname, email) VALUES (
        $1, $2, $3, $4
    )
    RETURNING *;

-- name: ListSwipableProfiles :many
SELECT * FROM users WHERE users