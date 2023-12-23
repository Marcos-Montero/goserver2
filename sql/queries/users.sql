-- name: CreateUser :one

INSERT INTO USERS (
    ID,
    CREATED_AT,
    UPDATED_AT,
    NAME,
    API_KEY
) VALUES (
    $1,
    $2,
    $3,
    $4,
    ENCODE(SHA256(RANDOM()::TEXT::BYTEA), 'hex')
) RETURNING *;

-- name: GetUserByAPIKey :one
SELECT
    *
FROM
    USERS
WHERE
    API_KEY = $1;