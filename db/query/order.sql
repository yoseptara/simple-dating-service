-- name: CreateOrder :one
INSERT INTO
    orders (
        customer_email,
        xendit_invoice_id,
        payment_status,
        esim_id,
        country_code,
        plan_option,
        data_amount,
        data_unit,
        duration_in_days,
        option_id,
        idr_price,
        quantity
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
    )
RETURNING *;

-- name: AddOrderTopUpId :one
INSERT INTO
    order_topup_ids (
        order_id,
        usimsa_topup_id
    ) VALUES (
        $1, $2
    )
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetOrderByTopupId :one
SELECT * FROM orders
INNER JOIN order_topup_ids
ON orders.id = order_topup_ids.order_id
WHERE order_topup_ids.usimsa_topup_id = $1 
LIMIT 1;

-- name: UpdateOrder :one
UPDATE orders SET payment_status = $3,
paid_at = $4
WHERE unique_order_id = $1
AND xendit_invoice_id = $2
RETURNING *;

-- name: AddOrderInvoice :one
UPDATE orders SET xendit_invoice_id = $2,
unique_order_id = $3,
payment_status = $4
WHERE id = $1
RETURNING *;