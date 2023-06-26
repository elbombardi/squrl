-- name: CheckShortUrlKeyExists :one
SELECT EXISTS(SELECT 1 FROM short_urls WHERE short_url_key = $1 AND customer_id = $2);

-- name: InsertNewShortURL :exec
INSERT INTO short_urls (short_url_key, customer_id, long_url)
VALUES ($1, $2, $3)
RETURNING id, short_url_key, customer_id, long_url, status, tracking_status, click_count, first_click_date_time, last_click_date_time, created_at, updated_at;

-- name: UpdateShortURLStatus :exec
UPDATE short_urls SET status = $1, updated_at=now() WHERE short_url_key = $2 AND customer_id = $3
RETURNING id, short_url_key, customer_id, long_url, status, tracking_status, click_count, first_click_date_time, last_click_date_time, created_at, updated_at;

-- name: UpdateShortURLTrackingStatus :exec
UPDATE short_urls SET tracking_status = $1, updated_at=now() WHERE short_url_key = $2 AND customer_id = $3
RETURNING id, short_url_key, customer_id, long_url, status, tracking_status, click_count, first_click_date_time, last_click_date_time, created_at, updated_at;

-- name: UpdateShortURLLongURL :exec
UPDATE short_urls SET long_url = $1, updated_at=now() WHERE short_url_key = $2 AND customer_id = $3
RETURNING id, short_url_key, customer_id, long_url, status, tracking_status, click_count, first_click_date_time, last_click_date_time, created_at, updated_at;

-- name: IncrementShortURLClickCount :exec
UPDATE short_urls SET click_count = click_count + 1, updated_at=now() WHERE short_url_key = $1 AND customer_id = $2
RETURNING id, short_url_key, customer_id, long_url, status, tracking_status, click_count, first_click_date_time, last_click_date_time, created_at, updated_at;

-- name: SetShortURLFirstClickDate :exec
UPDATE short_urls SET first_click_date_time = $1, updated_at=now() WHERE short_url_key = $2 AND customer_id = $3
RETURNING id, short_url_key, customer_id, long_url, status, tracking_status, click_count, first_click_date_time, last_click_date_time, created_at, updated_at;

-- name: SetShortURLLastClickDate :exec
UPDATE short_urls SET last_click_date_time = $1, updated_at=now() WHERE short_url_key = $2 AND customer_id = $3
RETURNING id, short_url_key, customer_id, long_url, status, tracking_status, click_count, first_click_date_time, last_click_date_time, created_at, updated_at;

-- name: GetShortURLByCustomerIDAndShortURLKey :one
SELECT id, short_url_key, customer_id, long_url, status, tracking_status, click_count, first_click_date_time, last_click_date_time, created_at, updated_at
FROM short_urls WHERE customer_id = $1 AND short_url_key = $2;