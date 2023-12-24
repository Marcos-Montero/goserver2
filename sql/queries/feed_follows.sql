-- name: CreateFeedFollow :one
INSERT INTO FEED_FOLLOWS(
    ID,
    CREATED_AT,
    UPDATED_AT,
    USER_ID,
    FEED_ID
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: GetFeedFollows :many
SELECT
    *
FROM
    FEED_FOLLOWS
WHERE
    USER_ID = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM FEED_FOLLOWS
WHERE
    ID = $1
    AND USER_ID = $2;