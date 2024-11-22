-- +goose Up
-- +goose StatementBegin
CREATE TABLE standarts(
    id SERIAL PRIMARY KEY,
    teacher_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    standart_id INT NOT NULL,
    subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    type VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS standarts CASCADE;
-- +goose StatementEnd
