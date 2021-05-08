package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type ReviewRepository interface {
	FindAllReview() (entity.Reviews, error)
	FindReviewByID(id int) (entity.Review, error)
	CreateReview(book entity.Review) error
	UpdateReview(book entity.Review) error
	DeleteReviewByID(id int) error
}
