-- +goose Up
create table users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR NOT NULL UNIQUE
);

-- +goose Down
drop table users;