package loaders

import (
	"context"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/samber/do"
	lop "github.com/samber/lo/parallel"
	"github.com/t-kuni/go-graphql-template/domain/infrastructure/db"
	"github.com/t-kuni/go-graphql-template/ent"
	"github.com/t-kuni/go-graphql-template/ent/user"
	"github.com/t-kuni/go-graphql-template/graph/model"
)

// LoaderImpl reads Users from a database
type LoaderImpl struct {
	conn db.Connector
}

func NewLoaderImpl(i *do.Injector) (*LoaderImpl, error) {
	return &LoaderImpl{
		do.MustInvoke[db.Connector](i),
	}, nil
}

// GetUsers implements a batch function that can retrieve many users by ID,
// for use in a dataloader
func (l *LoaderImpl) GetUsers(ctx context.Context, userIds []int) []*dataloader.Result[*model.User] {
	users, err := l.conn.GetEnt().User.Query().Where(user.IDIn(userIds...)).All(ctx)
	if err != nil {
		return handleError[*model.User](len(userIds), err)
	}

	idModelMap := make(map[int]*ent.User)
	for _, m := range users {
		idModelMap[m.ID] = m
	}

	return lop.Map(userIds, func(id int, _ int) *dataloader.Result[*model.User] {
		m := idModelMap[id]
		return &dataloader.Result[*model.User]{Data: &model.User{ID: m.ID, Name: m.Name, Age: m.Age}}
	})
}

// handleError creates array of result with the same error repeated for as many items requested
func handleError[T any](itemsLength int, err error) []*dataloader.Result[T] {
	result := make([]*dataloader.Result[T], itemsLength)
	for i := 0; i < itemsLength; i++ {
		result[i] = &dataloader.Result[T]{Error: err}
	}
	return result
}
