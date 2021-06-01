package registry

import (
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/mysql"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/adapter"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/controller"
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
		{
			Name:  "auth-controller",
			Build: buildAuthenticationController,
		},
	}...); err != nil {
		return nil, err
	}
	return &Container{
		ctn: builder.Build(),
	}, nil
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func (c *Container) Clean() error {
	return c.ctn.Clean()
}

func buildBookShelfController(ctn di.Container) (interface{}, error) {
	conn := mysql.GetMySQLConnection()
	handler := mysql.NewMySQLHandler(conn)
	adapter.NewBookRepository(handler)
	adapter.NewCommentRepository(handler)
	adapter.NewReviewRepository(handler)
	adapter.NewShelfRepository(handler)
	adapter.NewUserRepository(handler)

	bookInputport := interactor.NewBookInteractor()
	commentInputport := interactor.NewCommentInteractor()
	reviewInputport := interactor.NewReviewInteractor()
	shelfInputport := interactor.NewShelfInteractor()
	userInputport := interactor.NewUserInteractor()
	ctr := controller.NewBookshelfController(
		bookInputport,
		commentInputport,
		reviewInputport,
		shelfInputport,
		userInputport,
	)
	return &ctr, nil
}

func buildAuthenticationController(ctn di.Container) (interface{}, error) {
	userInputport := interactor.NewUserInteractor()
	ctr := controller.NewAuthController(userInputport)
	return &ctr, nil
}
