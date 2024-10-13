-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *
)
SELECT
inserted_feed_follow.*,
users.name AS user_name,
feeds.name AS feed_name
FROM inserted_feed_follow
Inner JOIN users on users.id = inserted_feed_follow.user_id
Inner JOIN feeds on feeds.id = inserted_feed_follow.feed_id;
--where inserted_feed_follow.id = $1;

-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follows.id,
    feed_follows.created_at,
    feed_follows.updated_at,
    users.name AS user_name,
    feeds.name AS feed_name
FROM feed_follows
Inner JOIN users on users.id = feed_follows.user_id
inner JOIN feeds on feeds.id = feed_follows.feed_id
Where feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_id = $1
AND user_id = $2;