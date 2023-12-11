package repository

import (
	db "auth-server/db/sqlc"
	"context"
)

type Authorization interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUser(ctx context.Context, email string) (db.User, error)
	CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error)
	GetSession(ctx context.Context, id uuid.UUID) (db.Session, error)
}

type Repository struct {
	Authorization
}

func NewRepository(store Store) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(store),
	}
}
