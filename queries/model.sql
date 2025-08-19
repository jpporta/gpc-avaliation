-- name: CreateModel :one
INSERT INTO model (name, author, description, category_id)
VALUES (?, ?, ?, ?) RETURNING id;
