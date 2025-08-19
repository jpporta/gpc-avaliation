-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(255) NOT NULL UNIQUE,
	description TEXT,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE model
ADD COLUMN category_id INT NOT NULL REFERENCES category(id) ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS juror_categories(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	juror_id INT NOT NULL REFERENCES juror(id) ON DELETE CASCADE,
	category_id INT NOT NULL REFERENCES category(id) ON DELETE CASCADE,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	UNIQUE (juror_id, category_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS juror_categories;
ALTER TABLE model
DROP COLUMN category_id;
DROP TABLE IF EXISTS category;
-- +goose StatementEnd
