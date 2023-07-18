-- +goose Up
-- +goose StatementBegin
create table users
(
    id            uuid primary key,
    created_at    timestamptz default now(),
    username      text not null check ( char_length(username) >= 1 AND char_length(username) <= 32),
    password_hash text not null,
    unique (username)
);

create unique index idx_username on users (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index idx_username;
drop table users;
-- +goose StatementEnd