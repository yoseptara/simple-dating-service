-- name: CreateOrUpdateSwipe :one
INSERT INTO swipes (user_id, target_id, direction, swipe_date)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
ON CONFLICT (user_id, target_id) DO UPDATE
SET direction = $3, swipe_date = CURRENT_TIMESTAMP
RETURNING *;
