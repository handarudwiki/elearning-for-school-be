-- +goose Up
-- +goose StatementBegin
CREATE TABLE questions(
    id SERIAL PRIMARY KEY,
    question TEXT NOT NULL,
    question_type INT NOT NULL,
    question_bank_id INT NOT NULL REFERENCES question_banks(id) ON DELETE CASCADE,
    audio TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists questions CASCADE;
-- +goose StatementEnd
