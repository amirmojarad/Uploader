-- +goose Up
create table files (
                       id serial primary key,
                       user_id bigint not null,
                       bucket_id bigint not null,
                       name varchar(255) not null ,
                       size int8 not null,
                       type varchar(255) not null,
                       created_at timestamptz not null default current_timestamp,
                       updated_at timestamptz not null default current_timestamp,
                       deleted_at timestamptz default null,
                       constraint fk_user_id foreign key (user_id) references users(id),
                       constraint fk_bucket_id foreign key (bucket_id) references buckets(id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table files;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
