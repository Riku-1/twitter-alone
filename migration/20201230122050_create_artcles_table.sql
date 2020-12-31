-- +goose Up
-- +goose StatementBegin
CREATE TABLE articles (
    id int NOT NULL primary key,
    title varchar(255) NOT NULL,
    body text NOT NULL,
    publish_date datetime,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE articles;
-- +goose StatementEnd