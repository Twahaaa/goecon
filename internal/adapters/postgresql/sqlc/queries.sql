-- name: ListProducts :many
SELECT *
FROM products;
-- name: FindProductsById :one
SELECT *
FROM products
WHERE id = $1;
-- name: ListOrders :many
SELECT *
FROM orders;
-- name: FindOrderById :many
SELECT orders.id,
    orders.customer_id,
    orders.created_at,
    order_items.product_id,
    order_items.quantity,
    order_items.price_cents
FROM orders
    INNER JOIN order_items ON orders.id = order_items.order_id
WHERE orders.id = $1;