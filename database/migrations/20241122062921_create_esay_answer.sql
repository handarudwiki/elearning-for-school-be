-- +goose Up
-- +goose StatementBegin
CREATE TABLE esay_answers(
    id SERIAL PRIMARY KEY,
    exam_schedule_id INT NOT NULL REFERENCES exam_schedules(id) ON DELETE CASCADE,
    student_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    answer_id INT NOT NULL REFERENCES student_answers(id) ON DELETE CASCADE,
    corrected_by INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    point INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS esay_answers CASCADE;
-- +goose StatementEnd
