package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type commentInteractor struct {
}

func NewCommentInteractor() inputport.CommentInputPort {
	return &commentInteractor{}
}

func (c *commentInteractor) FindAllComment() (entity.Comments, error) {
	panic("implement me")
}

func (c *commentInteractor) FindCommentByID(id int) (*entity.Comment, error) {
	panic("implement me")
}

func (c *commentInteractor) CreateComment(userID int, book entity.Comment) (*entity.Comment, error) {
	panic("implement me")
}

func (c *commentInteractor) UpdateComment(userID int, book entity.Comment) error {
	panic("implement me")
}

func (c *commentInteractor) DeleteCommentByID(userID int, id int) error {
	panic("implement me")
}
