
-- name: CheckUsernameExists :one
SELECT EXISTS(SELECT 1 FROM account WHERE username = $1);

-- name: CheckPrefixExists :one
SELECT EXISTS(SELECT 1 FROM account WHERE prefix = $1);

-- name: InsertNewAccount :exec
INSERT INTO account (prefix, username, email, hashed_password)
VALUES ($1, $2, $3, $4)
RETURNING id, prefix, username, email, hashed_password, enabled, created_at, updated_at;

-- name: UpdateAccountStatusByUsername :exec
UPDATE account SET enabled = $1, updated_at=now() WHERE username = $2
RETURNING id, prefix, username, email, hashed_password, enabled, created_at, updated_at;

-- name: GetAccountByUsername :one
SELECT id, prefix, username, email, hashed_password, enabled, created_at, updated_at
FROM account WHERE username = $1;

-- name: GetAccountByPrefix :one
SELECT id, prefix, username, email, hashed_password, enabled, created_at, updated_at
FROM account WHERE prefix = $1;
