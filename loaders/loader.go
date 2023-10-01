package loaders

import (
	"github.com/graph-gophers/dataloader/v7"
	"github.com/samber/do"
	"github.com/t-kuni/go-graphql-template/graph/model"
	"time"
)

// Loaders reads Users from a database
type Loaders struct {
	UserLoader *dataloader.Loader[int, *model.User]
}

func NewLoader(i *do.Injector) (*Loaders, error) {
	impl := do.MustInvoke[*LoaderImpl](i)

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader(impl.GetUsers, dataloader.WithWait[int, *model.User](time.Millisecond)),
	}, nil
}
