package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"

	"github.com/shintaro-uchiyama/go-ucwork/graph/generated"
	"github.com/shintaro-uchiyama/go-ucwork/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := domain.NewUser(input.Email, input.Password)
	err := r.userApplicationService.Create(*user)
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
