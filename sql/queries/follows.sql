-- name: CreateFeedFollow :one
WITH
    inserted_ff as (
        INSERT INTO
            follows (
                id,
                created_at,
                updated_at,
                user_id,
                feed_id
            )
        VALUES ($1, $2, $3, $4, $5) RETURNING *
    )
select
    inserted_ff.*,
    u.name AS username,
    f.name AS feed_name
FROM
    inserted_ff
    INNER JOIN users u ON u.id = inserted_ff.user_id
    INNER JOIN feeds f ON inserted_ff.feed_id = f.id;

-- name: GetFeedFollow :many
SELECT fe.* from follows fo inner join feeds fe on fo.feed_id = fe.id
WHERE fo.user_id = (SELECT id from users u where u.name = $1); 

-- name: RemoveFeedFollow :exec
delete from follows 
where
 user_id = (select id from users where users.name = $1)
 AND
 feed_id = (select id from feeds where feeds.url = $2);