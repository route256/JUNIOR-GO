-- +goose Up
-- +goose StatementBegin
create table articles(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name text NOT NULL DEFAULT '',
    rating int not null default 0,
    created_at timestamp with time zone default now() not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table articles;
-- +goose StatementEnd
