-- name: CheckShortUrlKeyExists :one
SELECT EXISTS(SELECT 1 FROM link WHERE short_url_key = $1 AND account_id = $2);

-- name: InsertNewLink :exec
INSERT INTO link (short_url_key, account_id, long_url)
VALUES ($1, $2, $3)
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at;

-- name: UpdateLinkStatus :exec
UPDATE link SET enabled = $1, updated_at=now() WHERE id = $2 
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at;

-- name: UpdateLinkTrackingStatus :exec
UPDATE link SET tracking_enabled = $1, updated_at=now() WHERE id = $2 
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at;

-- name: UpdateLinkLongURL :exec
UPDATE link SET long_url = $1, updated_at=now() WHERE id = $2
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at;

-- name: GetLinkByAccountIDAndShortURLKey :one
SELECT id, short_url_key, account_id, long_url, enabled, tracking_enabled,created_at, updated_at
FROM link WHERE account_id = $1 AND short_url_key = $2;