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
	return `INSERT INTO reviews (book_id, content, reading_date, create_user_id) VALUES (?, ?, ?, ?)`
}

func (r *reviewRepository) CreateReview(userID int, review entity.Review) (insReview entity.Review, err error) {
	query := r.getCreateReviewQuery()
	result, err := r.handler.Exec(query,
		review.BookID,
		review.Content,
		review.ReadingDate,
		userID,
	)
	if err != nil {
		return
	}
	insReview = review
	insID, err := result.LastInsertId()
	if err != nil {
		return
	}
	insReview.ID = int(insID)
	return insReview, nil
}

func (r *reviewRepository) getUpdateReviewQuery() string {
	return `UPDATE reviews SET
book_id = ?,
content = ?,
reading_date = ?,
create_user_id = ?
WHERE id = ?`
}

func (r *reviewRepository) UpdateReview(userID int, review entity.Review) error {
	query := r.getUpdateReviewQuery()
	_, err := r.handler.Exec(query,
		review.BookID,
		review.Content,
		review.ReadingDate,
		userID,
		review.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *reviewRepository) getDeleteReviewByIDQuery() string {
	return `UPDATE reviews
SET delete_user_id = ?,
delete_date = now(),
delete_flag = 1
WHERE id = ?`
}

func (r *reviewRepository) DeleteReviewByID(userID int, id int) error {
	query := r.getDeleteReviewByIDQuery()
	_, err := r.handler.Exec(query,
		userID,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
