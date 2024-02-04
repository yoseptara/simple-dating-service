-- name: CreateCountry :exec
INSERT INTO
    countries (code, name)
VALUES (
    $1, $2
)
ON CONFLICT (code)
DO NOTHING;

-- name: ListCountriesWithPrice :many
WITH MinPrices AS (
    SELECT 
        country_code, 
        MIN(idr_price) AS cheapest_idr_price
    FROM 
        esims
    GROUP BY 
        country_code
)

SELECT 
    c.code AS country_code,
    c.name AS country_name, 
    e.plan_option, 
    e.data_amount, 
    e.data_unit, 
    e.duration_in_days, 
    e.option_id, 
    e.idr_price
FROM 
    MinPrices m
JOIN 
    countries c ON m.country_code = c.code
JOIN 
    esims e ON m.country_code = e.country_code AND m.cheapest_idr_price = e.idr_price;
