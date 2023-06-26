-- name: InsertNewClick :exec
INSERT INTO clicks (id, short_url_id, click_date_time, user_agent, ip_address)
VALUES ($1, $2, $3, $4, $5);