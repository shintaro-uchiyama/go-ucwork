package infrastructure

import (
	"context"
	"fmt"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"

	"firebase.google.com/go/auth"

	firebase "firebase.google.com/go"
)

type UserFirebaseAuth struct {
	client *auth.Client
}

func NewUserFirebaseAuth() (*UserFirebaseAuth, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("[infrastructure.NewUserFirebaseAuth] initializing firebase app error: %w", err)
	}
	client, authErr := app.Auth(context.Background())
	if authErr != nil {
		return nil, fmt.Errorf("[infrastructure.NewUserFirebaseAuth] initializing firebase auth error: %w", authErr)
	}
	return &UserFirebaseAuth{
		client: client,
	}, nil
}

func (u *UserFirebaseAuth) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	params := (&auth.UserToCreate{}).
		Email(user.Email()).
		EmailVerified(false).
		Password(user.Password()).
		DisplayName(fmt.Sprintf("%w %w", user.FirstName(), user.LastName())).
		Disabled(false)
	createdUser, err := u.client.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("[infrastructure.CreateUser] error creating user: %w", err)
	}
	var domainUser domain.User
	domainUser.SetUUID(createdUser.UID)
	return &domainUser, nil
}

func (u *UserFirebaseAuth) GetUserByEmail(ctx context.Context, user *domain.User) (*domain.User, error) {
	foundUser, err := u.client.GetUserByEmail(ctx, user.Email())
	if err != nil {
		return nil, fmt.Errorf("[infrastructure.GetUserByEmail]: %w", err)
	}

	if foundUser != nil {
		return nil, nil
	}

	domainUser := domain.User{}
	domainUser.SetEmail(foundUser.Email)
	return &domainUser, nil
}
