-- +goose Up
-- +goose StatementBegin
CREATE TABLE classroom_lives(
    id SERIAL PRIMARY KEY,
    schedule_id INT NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    body TEXT NOT NULL,
    settings TEXT NOT NULL,
    is_active BOOLEAN DEFAULT false,
    note TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists classroom_lives CASCADE;
-- +goose StatementEnd
