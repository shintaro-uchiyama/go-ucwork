package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/shintaro-uchiyama/go-ucwork/graph/generated"
	"github.com/shintaro-uchiyama/go-ucwork/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("[graph.CreateTodo]: %w", err)
	}
	uu := u.String()
	user := &model.User{
		ID:   uu,
		UUID: uu,
		Name: input.Name,
	}
	r.users = append(r.users, user)
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
