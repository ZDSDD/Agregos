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