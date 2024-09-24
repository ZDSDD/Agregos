-- +goose Up
create table follows(
    id UUID primary key,
    created_at TIMESTAMP not null,
    updated_at TIMESTAMP not null,
    user_id UUID NOT NULL,
    feed_id UUID NOT NULL,
    constraint fk_user foreign key (user_id) references users(id) on delete cascade,
    constraint fk_feed foreign key (feed_id) references feeds(id) on delete cascade,
    constraint u_uf UNIQUE(user_id, feed_id)
);

-- +goose Down
drop table follows;