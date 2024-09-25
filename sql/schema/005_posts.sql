-- +goose Up
create table posts(
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title varchar not null,
    url varchar unique not null,
    description varchar,
    published_at timestamp,
    feed_id UUID not null,
    constraint feed_fk foreign key (feed_id) references feeds(id) on delete cascade
);

-- +goose Down
drop table posts;