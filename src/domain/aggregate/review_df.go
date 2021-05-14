package aggregate

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type ReviewDfs []*ReviewDf

type ReviewDf struct {
	Review  ReviewDfs
	Comment entity.Comments
}
