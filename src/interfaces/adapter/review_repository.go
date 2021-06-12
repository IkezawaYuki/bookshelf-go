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
	return `select id, book_id, title, content, reading_date from reviews where delete_flag = 0`
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
			&review.Title,
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
	return `select id, book_id, title, content, reading_date from reviews where id = ? and delete_flag = 0`
}

func (r *reviewRepository) FindReviewByID(id int) (*entity.Review, error) {
	query := r.getFindReviewByIDQuery()
	row := r.handler.QueryRow(query, id)
	var review entity.Review
	err := row.Scan(
		&review.ID,
		&review.BookID,
		&review.Title,
		&review.Content,
		&review.ReadingDate,
	)
	return &review, err
}

func (r *reviewRepository) getCreateReviewQuery() string {
	return `INSERT INTO reviews (book_id, title, content, reading_date, create_user_id) VALUES (?, ?, ?, ?)`
}

func (r *reviewRepository) CreateReview(userID int, review entity.Review) (*entity.Review, error) {
	query := r.getCreateReviewQuery()
	result, err := r.handler.Exec(query,
		review.BookID,
		review.Title,
		review.Content,
		review.ReadingDate,
		userID,
	)
	if err != nil {
		return nil, err
	}
	insID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	review.ID = int(insID)
	return &review, nil
}

func (r *reviewRepository) getUpdateReviewQuery() string {
	return `UPDATE reviews SET
book_id = ?,
title = ?,
content = ?,
reading_date = ?,
create_user_id = ?
WHERE id = ?`
}

func (r *reviewRepository) UpdateReview(userID int, review entity.Review) error {
	query := r.getUpdateReviewQuery()
	_, err := r.handler.Exec(query,
		review.BookID,
		review.Title,
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

func (r *reviewRepository) getFindByBookIDQuery() string {
	return `select 
id,
book_id,
title,
content,
reading_date
from reviews
where book_id = ?`
}

func (r *reviewRepository) FindReviewByBookID(id int) (entity.Reviews, error) {
	result := make(entity.Reviews, 0)
	query := r.getFindByBookIDQuery()
	rows, err := r.handler.Query(query, id)
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
			&review.Title,
			&review.Content,
			&review.ReadingDate,
		); err != nil {
			return nil, err
		}
		result = append(result, &review)
	}
	return result, nil
}
