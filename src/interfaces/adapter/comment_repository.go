package adapter

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
)

func NewCommentRepository(handler datastore.DBHandler) {
	CommentRepo = &commentRepository{handler: handler}
}

var CommentRepo repository.CommentRepository

type commentRepository struct {
	handler datastore.DBHandler
}

func (b *commentRepository) getFindAllCommentQuery() string {
	return `select id, review_id, content from comments where delete_flag = 0`
}

func (b *commentRepository) FindAllComment() (entity.Comments, error) {
	result := make(entity.Comments, 0)
	query := b.getFindAllCommentQuery()
	rows, err := b.handler.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		comment := entity.Comment{}
		if err := rows.Scan(
			&comment.ID,
			&comment.ReviewID,
			&comment.Content,
		); err != nil {
			return nil, err
		}
		result = append(result, &comment)
	}
	return result, nil
}

func (b *commentRepository) getFindCommentByIDQuery() string {
	return `select id, review_id, content from comments where id = ? and delete_flag = 0`
}

func (b *commentRepository) FindCommentByID(id int) (comment entity.Comment, err error) {
	query := b.getFindCommentByIDQuery()
	row := b.handler.QueryRow(query, id)
	err = row.Scan(
		&comment.ID,
		&comment.ReviewID,
		&comment.Content,
	)
	return
}

func (b *commentRepository) getCreateCommentQuery() string {
	return `INSERT INTO comments (review_id, content, create_user_id) VALUES (?, ?, ?);`
}

func (b *commentRepository) CreateComment(userID int, comment entity.Comment) (insComment entity.Comment, err error) {
	query := b.getCreateCommentQuery()
	result, err := b.handler.Exec(query,
		comment.ReviewID,
		comment.Content,
		userID,
	)
	if err != nil {
		return
	}
	insComment = comment
	insID, err := result.LastInsertId()
	if err != nil {
		return
	}
	insComment.ID = int(insID)
	return
}

func (b *commentRepository) getUpdateCommentQuery() string {
	return `UPDATE comments SET 
review_id = ?,
content = ?,
create_user_id = ?
where id = ?`
}

func (b *commentRepository) UpdateComment(userID int, comment entity.Comment) error {
	query := b.getUpdateCommentQuery()
	_, err := b.handler.Exec(query,
		comment.ReviewID,
		comment.Content,
		userID,
		comment.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (b *commentRepository) getDeleteCommentByIDQuery() string {
	return `UPDATE comments
SET delete_user_id = ?,
delete_date = now(),
delete_flag = 1
WHERE id = ?`
}

func (b *commentRepository) DeleteCommentByID(userID int, id int) error {
	query := b.getDeleteCommentByIDQuery()
	_, err := b.handler.Exec(query,
		userID,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (b *commentRepository) getFindCommentByReviewID() string {
	return `select id, review_id, content from comments where delete_flag = 0 and review_id = ?`
}

func (b *commentRepository) FindCommentByReviewID(reviewID int) (entity.Comments, error) {
	result := make(entity.Comments, 0)
	query := b.getFindCommentByReviewID()
	rows, err := b.handler.Query(query, reviewID)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		comment := entity.Comment{}
		if err := rows.Scan(
			&comment.ID,
			&comment.ReviewID,
			&comment.Content,
		); err != nil {
			return nil, err
		}
		result = append(result, &comment)
	}
	return result, nil
}
