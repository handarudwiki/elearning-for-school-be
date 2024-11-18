-- +goose Up
-- +goose StatementBegin
CREATE TABLE classroom_tasks (
    id SERIAL PRIMARY KEY,
    classroom_id INT NOT NULL REFERENCES classrooms(id) ON DELETE CASCADE,
    task_id INT NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    teacher_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    body TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS classrooms_tasks CASCADE;
-- +goose StatementEnd
