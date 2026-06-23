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

-- name: CreateProduct :one
INSERT INTO products (name, price_in_cents, quantity) VALUES ($1, $2, $3) RETURNING *;

-- name: CreateOrder :one
INSERT INTO orders (customer_id) VALUES ($1) RETURNING *;

-- name: FetchPrice :one
SELECT price_in_cents FROM products WHERE id = ($1);

-- name: CreateOrderItems :one
INSERT INTO order_items (order_id, product_id, quantity, price_cents) VALUES ($1, $2, $3, $4) RETURNING *;