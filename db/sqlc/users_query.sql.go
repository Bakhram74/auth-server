// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: users_query.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" ("username",
                   "email",
                   "password")
VALUES ($1, $2, $3) RETURNING id, username, email, password, "isActivated", "isActivationLink"
`

type CreateUserParams struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Username, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.IsActivated,
		&i.IsActivationLink,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, password, "isActivated", "isActivationLink" FROM "user"
WHERE "email" = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.IsActivated,
		&i.IsActivationLink,
	)
	return i, err
}
