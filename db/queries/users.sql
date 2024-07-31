-- name: CreateUser :one
INSERT INTO users (
  username, password, display_name, email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- -- name: UpdateAuthor :exec
-- UPDATE authors
--   set name = $2,
--   bio = $3
-- WHERE id = $1;

-- -- name: DeleteAuthor :exec
-- DELETE FROM authors
-- WHERE id = $1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: FindUserByFilters :many
SELECT * FROM users
WHERE
  (username = $1 OR $1 IS NULL) AND
  (display_name = $2 OR $2 IS NULL) AND
  (id = $3 OR $3 IS NULL);

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;