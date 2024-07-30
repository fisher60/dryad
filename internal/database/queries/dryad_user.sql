-- name: GetDryadUser :one
SELECT * FROM dryad_user
WHERE id = $1 LIMIT 1;

-- name: ListDryadUsers :many
SELECT * FROM dryad_user;

-- name: CreateDryadUser :one
INSERT INTO dryad_user (abandonauth_uuid) VALUES ($1)
RETURNING *;
