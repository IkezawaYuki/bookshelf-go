package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type CommentRepository interface {
	FindAllComment() (entity.Comments, error)
	FindCommentByID(id int) (entity.Comment, error)
	CreateComment(book entity.Comment) error
	UpdateComment(book entity.Comment) error
	DeleteCommentByID(id int) error
}
