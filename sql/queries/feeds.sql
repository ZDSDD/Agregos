-- name: CreateFeed :one
insert into
    feeds (
        id,
        created_at,
        updated_at,
        name,
        url,
        user_id
    )
values ($1, $2, $3, $4, $5, $6) returning *;

-- name: GetFeeds :many
select f.name, f.url, u.name from feeds f inner join users u on user_id = u_id;