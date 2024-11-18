-- +goose Up
-- +goose StatementBegin
CREATE TABLE classroom_subjects(
    id SERIAL PRIMARY KEY,
    classroom_id INT NOT NULL REFERENCES classrooms(id) ON DELETE CASCADE,
    subject_id INT NOT NULL REFERENCES subjects(id) ON DELETE CASCADE,
    teacher_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists classroom_subject CASCADE;
-- +goose StatementEnd
