-- +goose Up
-- +goose StatementBegin
CREATE TABLE exam_schedules(
    id SERIAL PRIMARY KEY,
    question_bank_id INT NOT NULL REFERENCES question_banks(id) ON DELETE CASCADE,
    teacher_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type INT DEFAULT 0, -- | 0-> NO EVENT | 1-> UH| 2->UTS | 3->UAS
    classrooms TEXT DEFAULT NULL,
    name VARCHAR(255) NOT NULL,
    date TIMESTAMP NOT NULL,
    start_time TIME NOT NULL,
    duration INT  NOT NULL,
    is_active INT NOT NULL,
    settings TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS exam_schedules CASCADE;
-- +goose StatementEnd
