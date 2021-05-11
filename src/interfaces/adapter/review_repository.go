package adapter

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
)

func NewReviewRepository(handler datastore.DBHandler) repository.ReviewRepository {
	return &reviewRepository{handler: handler}
}

type reviewRepository struct {
	handler datastore.DBHandler
}

func (r *reviewRepository) getFindAllReviewQuery() string {
	return `select id, book_id, content, reading_date from reviews where delete_flag = 0`
}

func (r *reviewRepository) FindAllReview() (entity.Reviews, error) {
	result := make(entity.Reviews, 0)
	query := r.getFindAllReviewQuery()
	rows, err := r.handler.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		review := entity.Review{}
		if err := rows.Scan(
			&review.ID,
			&review.BookID,
			&review.Content,
			&review.ReadingDate,
		); err != nil {
			return nil, err
		}
		result = append(result, &review)
	}
	return result, nil
}

func (r *reviewRepository) getFindReviewByIDQuery() string {
	return `select id, book_id, content, reading_date from reviews where id = ? and delete_flag = 0`
}

func (r *reviewRepository) FindReviewByID(id int) (review entity.Review, err error) {
	query := r.getFindReviewByIDQuery()
	row := r.handler.QueryRow(query, id)
	err = row.Scan(
		&review.ID,
		&review.BookID,
		&review.Content,
		&review.ReadingDate,
	)
	return
}

func (r *reviewRepository) getCreateReviewQuery() string {

}

func (r *reviewRepository) CreateReview(userID int, review entity.Review) error {
	panic("implement me")
}

func (r *reviewRepository) getUpdateReviewQuery() string {

}

func (r *reviewRepository) UpdateReview(userID int, review entity.Review) error {
	panic("implement me")
}

func (r *reviewRepository) getDeleteReviewByIDQuery() string {

}

func (r *reviewRepository) DeleteReviewByID(userID int, id int) error {
	panic("implement me")
}
