-- name: CreatePost :one

INSERT INTO POSTS (
    ID,
    CREATED_AT,
    UPDATED_AT,
    TITLE,
    DESCRIPTION,
    PUBLISHED_AT,
    URL,
    FEED_ID
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING *;

-- name: GetPostsForUser :many

SELECT
    POSTS.*
FROM
    POSTS
    JOIN FEED_FOLLOWS
    ON POSTS.FEED_ID = FEED_FOLLOWS.FEED_ID
WHERE
    FEED_FOLLOWS.USER_ID = $1
ORDER BY
    POSTS.PUBLISHED_AT DESC LIMIT $2;