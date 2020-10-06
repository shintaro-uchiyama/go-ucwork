package domain

import (
	"context"
)

type FirebaseAuthInterface interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, user *User) (*User, error)
}
