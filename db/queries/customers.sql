
-- name: CheckUsernameExists :one
SELECT EXISTS(SELECT 1 FROM customers WHERE username = $1);

-- name: CheckEmailExists :one
SELECT EXISTS(SELECT 1 FROM customers WHERE email = $1);

-- name: CheckApiKeyExists :one
SELECT EXISTS(SELECT 1 FROM customers WHERE api_key = $1);

-- name: InsertNewCustomer :exec
INSERT INTO customers (prefix, username, email, api_key)
VALUES ($1, $2, $3, $4)
RETURNING id, prefix, username, email, api_key, status, created_at, updated_at;

-- name: UpdateCustomerStatusByUsername :exec
UPDATE customers SET status = $1, updated_at=now() WHERE username = $2
RETURNING id, prefix, username, email, api_key, status, created_at, updated_at;

-- name: GetCustomerByUsername :one
SELECT id, prefix, username, email, api_key, status, created_at, updated_at
FROM customers WHERE username = $1;

-- name: GetCustomerByPrefix :one
SELECT id, prefix, username, email, api_key, status, created_at, updated_at
FROM customers WHERE prefix = $1;

-- name: GetCustomerByApiKey :one
SELECT id, prefix, username, email, api_key, status, created_at, updated_at
FROM customers WHERE api_key = $1;


