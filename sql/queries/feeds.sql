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
SELECT 
    f.name AS feed_name,
    f.url,
    u.name AS username
FROM feeds f
INNER JOIN users u ON f.user_id = u.id;