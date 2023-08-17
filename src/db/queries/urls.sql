-- name: CheckShortUrlKeyExists :one
SELECT EXISTS(SELECT 1 FROM url WHERE short_url_key = $1 AND account_id = $2);

-- name: InsertNewURL :exec
INSERT INTO url (short_url_key, account_id, long_url)
VALUES ($1, $2, $3)
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at;

-- name: UpdateURLStatus :exec
UPDATE url SET enabled = $1, updated_at=now() WHERE id = $2 
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at;

-- name: UpdateURLTrackingStatus :exec
UPDATE url SET tracking_enabled = $1, updated_at=now() WHERE id = $2 
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at;

-- name: UpdateLongURL :exec
UPDATE url SET long_url = $1, updated_at=now() WHERE id = $2
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at;

-- name: GetURLByAccountIDAndShortURLKey :one
SELECT id, short_url_key, account_id, long_url, enabled, tracking_enabled,created_at, updated_at
FROM url WHERE account_id = $1 AND short_url_key = $2;