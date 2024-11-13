-- +goose Up
-- +goose StatementBegin
CREATE TABLE lecture_comments (
    id SERIAL PRIMARY KEY,
    lecture_id INT NOT NULL REFERENCES lectures(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lecture_comments CASCADE;
-- +goose StatementEnd
