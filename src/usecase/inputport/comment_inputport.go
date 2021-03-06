package inputport

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type CommentInputPort interface {
	FindAllComment(page int, search string) (entity.Comments, error)
	FindCommentByID(id int) (*entity.Comment, error)
	CreateComment(userID int, book entity.Comment) (*entity.Comment, error)
	UpdateComment(userID int, book entity.Comment) error
	DeleteCommentByID(userID int, id int) error
	FindCommentByReviewID(id int) (entity.Comments, error)
}
