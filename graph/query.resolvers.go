package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/shintaro-uchiyama/go-ucwork/graph/generated"
	"github.com/shintaro-uchiyama/go-ucwork/graph/model"
)

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.userApplicationService.List()
	if err != nil {
		return nil, err
	}
	var responseUsers []*model.User
	for _, user := range users {
		responseUsers = append(responseUsers, &model.User{
			Email: user.Email(),
		})
	}
	return responseUsers, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
