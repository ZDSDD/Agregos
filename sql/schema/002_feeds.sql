-- +goose Up
create table feeds (
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    name VARCHAR not null,
    url VARCHAR UNIQUE not null,
    user_id UUID not null,
    constraint fk_user foreign key (user_id) references users (id) on delete cascade
);

-- +goose Down
drop table feeds;