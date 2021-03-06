package graph

import "github.com/shintaro-uchiyama/go-ucwork/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users                  []*model.User
	userApplicationService UserApplicationServiceInterface
}

func NewResolver(userApplicationService UserApplicationServiceInterface) *Resolver {
	return &Resolver{
		userApplicationService: userApplicationService,
	}
}
