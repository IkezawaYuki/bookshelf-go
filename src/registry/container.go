package registry

import "github.com/sarulabs/di"

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
	panic("implement me")
}
