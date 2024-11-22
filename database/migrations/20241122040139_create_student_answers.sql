-- +goose Up
-- +goose StatementBegin
CREATE table student_answers(
    id SERIAL PRIMARY KEY,
    question_bank_id INT NOT NULL REFERENCES question_banks(id) ON DELETE CASCADE,
    question_id INT NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    student_id  INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    exam_schedule_id INT NOT NULL REFERENCES exam_schedules(id) ON DELETE CASCADE,
    answer INT NULL,
    esay TEXT NULL,
    doubt BOOLEAN DEFAULT FALSE,
    is_correct BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS student_answers CASCADE;
-- +goose StatementEnd
