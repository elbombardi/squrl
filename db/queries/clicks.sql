-- name: InsertNewClick :exec
INSERT INTO clicks (short_url_id, user_agent, ip_address)
VALUES ($1, $2, $3);