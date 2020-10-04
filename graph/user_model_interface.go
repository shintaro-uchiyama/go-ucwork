package graph

import (
	"github.com/shintaro-uchiyama/go-ucwork/graph/model"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

var _ ModelUserConverter = (*model.NewUser)(nil)

type ModelUserConverter interface {
	ToDomainUser() *domain.User
}
