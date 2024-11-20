-- +goose Up
-- +goose StatementBegin
CREATE TABLE events(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    location VARCHAR(255) NOT NULL,
    title VARCHAR(100) NOT NULL,
    body TEXT NOT NULL,
    date DATE NOT NULL,
    time TIME NOT NULL,
    settings TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists events CASCADE;
-- +goose StatementEnd
