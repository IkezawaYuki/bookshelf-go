package aggregate

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type BookDfs []BookDf

type BookDf struct {
	Book    *entity.Book
	Reviews entity.Reviews
}
