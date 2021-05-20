package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type CommentRepository interface {
	FindAllComment() (entity.Comments, error)
	FindCommentByID(id int) (entity.Comment, error)
	CreateComment(userID int, book entity.Comment) (entity.Comment, error)
	UpdateComment(userID int, book entity.Comment) error
	DeleteCommentByID(userID int, id int) error
	FindCommentByReviewID(reviewID int) (entity.Comments, error)
}
