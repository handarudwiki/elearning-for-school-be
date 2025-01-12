-- +goose Up
-- +goose StatementBegin
CREATE TABLE abcents(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    schedule_id INT NOT NULL REFERENCES schedules(id) ON DELETE CASCADE,
    is_abcent BOOLEAN NOT NULL,
    reason INT NOT NULL,-- (0.No Reason, 1.Alpha, 2.Sick, 3.Permit, 4.Problem)
    description TEXT NULL,
    details TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE if exists abcents CASCADE;
-- +goose StatementEnd
