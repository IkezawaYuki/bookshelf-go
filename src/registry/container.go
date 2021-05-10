package registry

import (
	"fmt"
	"github.com/IkezawaYuki/bookshelf-go/src/infrastructure/mysql_client"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/adapter"
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
	conn := mysql_client.GetMySQLConnection()
	handler := mysql_client.NewMySQLHandler(conn)
	bookRepo := adapter.NewBookRepository(handler)
	commentRepo := adapter.NewCommentRepository(handler)
	fmt.Println(bookRepo)
	fmt.Println(commentRepo)

	panic("implement me")
}
