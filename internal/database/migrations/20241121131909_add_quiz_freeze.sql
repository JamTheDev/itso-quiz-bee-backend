-- +goose Up
-- +goose StatementBegin
ALTER TABLE quizzes
RENAME COLUMN is_active TO is_frozen;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
