-- +goose Up
create table users (
    id serial primary key,
    email varchar(255) not null unique,
    hashed_password varchar(255) not null,
    refresh_token varchar(255) not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    deleted_at timestamptz default null
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table users;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
