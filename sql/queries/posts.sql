-- name: CreatePost :exec
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $2, $3, $4, $5, $6);

-- name: GetPostsForUser :many
SELECT p.*
FROM posts p
JOIN feeds f ON p.feed_id = f.id
JOIN users u ON f.user_id = u.id
WHERE u.id = $1
ORDER BY p.published_at DESC
LIMIT $2;
