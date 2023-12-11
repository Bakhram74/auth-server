-- name: CreateUser :one
INSERT INTO "user" ("username",
                   "email",
                   "password")
VALUES ($1, $2, $3) RETURNING *;

-- name: GetUser :one
SELECT * FROM "user"
WHERE "email" = $1 LIMIT 1;




