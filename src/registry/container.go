package registry

import (
	"fmt"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/adapter"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/interactor"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

func NewContainer() (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	if err := builder.Add([]di.Def{
		{
			Name:  "bookshelf-controller",
			Build: buildBookShelfController,
		},
	}...); err != nil {
		return nil, err
	}
	return &Container{
		ctn: builder.Build(),
	}, nil
}

func buildBookShelfController(ctn di.Container) (interface{}, error) {
	conn := infrastructure.GetMySQLConnection()
	handler := infrastructure.NewMySQLHandler(conn)
	bookRepo := adapter.NewBookRepository(handler)
	commentRepo := adapter.NewCommentRepository(handler)
	reviewRepo := adapter.NewReviewRepository(handler)
	shelfRepo := adapter.NewShelfRepository(handler)
	userRepo := adapter.NewUserRepository(handler)
	itr := interactor.NewBookshelfInteractor(
		bookRepo,
		commentRepo,
		reviewRepo,
		shelfRepo,
		userRepo,
	)
	fmt.Println(itr)
	panic("")
}
