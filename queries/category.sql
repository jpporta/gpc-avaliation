-- name: CreateCategory :one
INSERT INTO category (name, description)
VALUES (?, ?) RETURNING id;

