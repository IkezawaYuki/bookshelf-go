package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type ReviewRepository interface {
	FindAllReview() (entity.Reviews, error)
	FindReviewByID(id int) (entity.Review, error)
	CreateReview(userID int, review entity.Review) (entity.Review, error)
	UpdateReview(userID int, review entity.Review) error
	DeleteReviewByID(userID int, id int) error
	FindByBookID(id int) (entity.Reviews, error)
}
