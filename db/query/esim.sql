-- name: CreateEsim :one
INSERT INTO
    esims (
        country_code,
        plan_option,
        data_amount,
        data_unit,
        duration_in_days,
        option_id,
        idr_price
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7
    )
    ON CONFLICT (option_id)
    DO UPDATE SET country_code = $1,
    plan_option=$2,
    data_amount= $3,
    data_unit =$4,
    duration_in_days=$5,
    idr_price=$7
    RETURNING *;

-- name: ListEsimsByCountry :many
SELECT esims.*, countries.name AS country_name
FROM esims
LEFT JOIN countries ON countries.code = esims.country_code
WHERE esims.country_code = $1;

-- name: GetEsim :one
SELECT * FROM esims
WHERE id = $1 LIMIT 1;

-- -- name: UpdateEsimStock :one
-- UPDATE esims SET stock= $2
-- where id = $1
-- RETURNING *;

-- name: DeleteEsim :exec
DELETE FROM esims WHERE id = $1;