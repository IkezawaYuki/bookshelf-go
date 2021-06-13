package registry

import (
	"database/sql"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/mysql"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/spreadsheet"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/adapter"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/controller"
	"github.com/IkezawaYuki/bookshelf-go/src/logger"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/interactor"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

var (
	conn *sql.DB
)

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
	if err := conn.Close(); err != nil {
		logger.Error("close is failed", err)
	}
	return c.ctn.Clean()
}

func buildBookShelfController(ctn di.Container) (interface{}, error) {
	conn = mysql.GetMySQLConnection()
	handler := mysql.NewMySQLHandler(conn)
	bookRepo := adapter.NewBookRepository(handler)
	commentRepo := adapter.NewCommentRepository(handler)
	reviewRepo := adapter.NewReviewRepository(handler)
	shelfRepo := adapter.NewShelfRepository(handler)
	userRepo := adapter.NewUserRepository(handler)

	bookInputport := interactor.NewBookInteractor(bookRepo)
	commentInputport := interactor.NewCommentInteractor(commentRepo)
	reviewInputport := interactor.NewReviewInteractor(reviewRepo)
	shelfInputport := interactor.NewShelfInteractor(shelfRepo)
	userInputport := interactor.NewUserInteractor(userRepo)
	presenter := outputport.NewPresenter()
	spreadsheetClient := spreadsheet.NewClient()
	ctr := controller.NewBookshelfController(
		bookInputport,
		commentInputport,
		reviewInputport,
		shelfInputport,
		userInputport,
		presenter,
		spreadsheetClient,
	)
	return &ctr, nil
}

func buildAuthenticationController(ctn di.Container) (interface{}, error) {
	conn := mysql.GetMySQLConnection()
	handler := mysql.NewMySQLHandler(conn)
	userRepo := adapter.NewUserRepository(handler)
	userInputport := interactor.NewUserInteractor(userRepo)
	ctr := controller.NewAuthController(userInputport)
	return &ctr, nil
}
