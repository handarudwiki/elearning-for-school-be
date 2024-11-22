-- +goose Up
-- +goose StatementBegin
CREATE TABLE question_banks(
    id SERIAL PRIMARY KEY,
    code VARCHAR(255) NOT NULL,
    standart_id INT DEFAULT 0,
    mc_count INT,--multiple choice count
    mc_option_count INT DEFAULT 4,
    esay_count INT DEFAULT 0,
    percentage TEXT,
    subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists question_banks CASCADE;
-- +goose StatementEnd
