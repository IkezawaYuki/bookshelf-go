package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type reviewInteractor struct {
	reviewRepo repository.ReviewRepository
}

func NewReviewInteractor(repo repository.ReviewRepository) inputport.ReviewInputPort {
	return &reviewInteractor{
		reviewRepo: repo,
	}
}

func (r *reviewInteractor) FindAllReview() (entity.Reviews, error) {
	return r.reviewRepo.FindAllReview()
}

func (r *reviewInteractor) FindReviewByID(id int) (*entity.Review, error) {
	return r.reviewRepo.FindReviewByID(id)
}

func (r *reviewInteractor) CreateReview(userID int, review entity.Review) (*entity.Review, error) {
	return r.reviewRepo.CreateReview(userID, review)
}

func (r *reviewInteractor) UpdateReview(userID int, review entity.Review) error {
	return r.reviewRepo.UpdateReview(userID, review)
}

func (r *reviewInteractor) DeleteReviewByID(userID int, id int) error {
	return r.reviewRepo.DeleteReviewByID(userID, id)
}
