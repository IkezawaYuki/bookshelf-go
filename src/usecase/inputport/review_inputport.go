package inputport

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type ReviewInputPort interface {
	FindAllReview() (entity.Reviews, error)
	FindReviewByID(id int) (*entity.Review, error)
	CreateReview(userID int, review entity.Review) (*entity.Review, error)
	UpdateReview(userID int, review entity.Review) error
	DeleteReviewByID(userID int, id int) error
}
