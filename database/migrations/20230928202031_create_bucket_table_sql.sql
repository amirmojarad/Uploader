-- +goose Up
create table buckets (
  id serial primary key,
  name varchar(255),
  user_id bigint,
  created_at timestamptz not null default current_timestamp,
  updated_at timestamptz not null default current_timestamp,
  deleted_at timestamptz default null,
  constraint fk_user_id foreign key (user_id) references users(id)
);
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
drop table buckets;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
