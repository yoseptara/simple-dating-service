-- name: CreateUser :one
INSERT INTO
    users (username, password, fullname, email) VALUES (
        $1, $2, $3, $4
    )
    RETURNING *;

-- name: ListSwipableProfiles :many
SELECT *
FROM users
LEFT JOIN swipes ON users.id = swipes.target_id
WHERE swipes.user_id != $1 OR (swipes.user_id = $1 AND swipes.swipe_date < CURRENT_TIMESTAMP - INTERVAL '1 day')
LIMIT $2;
