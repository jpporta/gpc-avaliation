-- name: CreateJuror :one
INSERT INTO juror (slug, name)
VALUES (?, ?) RETURNING id;

-- name: SetJurorCategories :exec
INSERT INTO juror_categories (juror_id, category_id)
VALUES (?, ?)
ON CONFLICT (juror_id, category_id) DO NOTHING;
