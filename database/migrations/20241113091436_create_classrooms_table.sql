-- +goose Up
-- +goose StatementBegin
CREATE TABLE classrooms(
    id SERIAL PRIMARY KEY,
    teacher_id INT  NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    grade VARCHAR(255) NOT NULL,
    "group" VARCHAR(255) NULL,
    settings VARCHAR(255) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS classrooms CASCADE;
-- +goose StatementEnd
