-- name: CreateFeed :one
INSERT INTO FEEDS (
    ID,
    CREATED_AT,
    UPDATED_AT,
    NAME,
    URL,
    USER_ID
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
) RETURNING *;

-- name: GetFeeds :many
SELECT
    *
FROM
    FEEDS;

-- name: GetNextFeedsToFetch :many

SELECT
    *
FROM
    FEEDS
ORDER BY
    LAST_FETCHED_AT ASC NULLS FIRST LIMIT $1;

-- name: MarkFeedAsFetched :one

UPDATE FEEDS
SET
    LAST_FETCHED_AT = NOW(
    ),
    UPDATED_AT = NOW(
    )
WHERE
    ID = $1 RETURNING *;