// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Session struct {
	ID           uuid.UUID `json:"id"`
	Userid       int32     `json:"userid"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	ID               int32       `json:"id"`
	Username         string      `json:"username"`
	Email            string      `json:"email"`
	Password         string      `json:"password"`
	IsActivated      pgtype.Bool `json:"isActivated"`
	IsActivationLink string      `json:"isActivationLink"`
}
