-- +goose Up
-- +goose StatementBegin
CREATE TABLE exam_results (
    id SERIAL PRIMARY KEY,
    exam_schedule_id INT NOT NULL REFERENCES exam_schedules(id) ON DELETE CASCADE,
    student_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    wrong_mc INT NOT NULL,
    wrong_esay INT NOT NULL,
    nul INT NOT NULL,
    result INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS exam_results CASCADE;
-- +goose StatementEnd
