package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type commentInteractor struct {
	commentRepo repository.CommentRepository
}

func NewCommentInteractor(repo repository.CommentRepository) inputport.CommentInputPort {
	return &commentInteractor{
		commentRepo: repo,
	}
}

func (c *commentInteractor) FindAllComment() (entity.Comments, error) {
	return c.commentRepo.FindAllComment()
}

func (c *commentInteractor) FindCommentByID(id int) (*entity.Comment, error) {
	return c.commentRepo.FindCommentByID(id)
}

func (c *commentInteractor) CreateComment(userID int, comment entity.Comment) (*entity.Comment, error) {
	return c.commentRepo.CreateComment(userID, comment)
}

func (c *commentInteractor) UpdateComment(userID int, comment entity.Comment) error {
	return c.commentRepo.UpdateComment(userID, comment)
}

func (c *commentInteractor) DeleteCommentByID(userID int, id int) error {
	return c.commentRepo.DeleteCommentByID(userID, id)
}

func (c *commentInteractor) FindCommentByReviewID(id int) (entity.Comments, error) {
	return c.commentRepo.FindCommentByReviewID(id)
}
