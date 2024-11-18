-- +goose Up
-- +goose StatementBegin
CREATE TABLE schedules(
    id SERIAL PRIMARY KEY,
    classroom_subject_id INT NOT NULL REFERENCES classroom_subjects(id) ON DELETE CASCADE,
    day INT NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists schedules CASCADE;
-- +goose StatementEnd
