-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    body TEXT NULL,
    type INT DEFAULT 0, -- 0: image, 1: FileUpload, 2: Essay
    is_active BOOLEAN DEFAULT FALSE,
    settings TEXT NULL,
    deadline TIMESTAMP NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks CASCADE;
-- +goose StatementEnd
