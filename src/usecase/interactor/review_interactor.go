package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type reviewInteractor struct {
}

func NewReviewInteractor() inputport.ReviewInputPort {
	return &reviewInteractor{}
}

func (r *reviewInteractor) FindAllReview() (entity.Reviews, error) {
	panic("implement me")
}

func (r *reviewInteractor) FindReviewByID(id int) (*entity.Review, error) {
	panic("implement me")
}

func (r *reviewInteractor) CreateReview(userID int, review entity.Review) (*entity.Review, error) {
	panic("implement me")
}

func (r *reviewInteractor) UpdateReview(userID int, review entity.Review) error {
	panic("implement me")
}

func (r *reviewInteractor) DeleteReviewByID(userID int, id int) error {
	panic("implement me")
}
