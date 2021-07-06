package adapter

import (
	"fmt"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/model"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
)

func NewCommentRepository(handler datastore.DBHandler) repository.CommentRepository {
	return &commentRepository{handler: handler}
}

type commentRepository struct {
	handler datastore.DBHandler
}

func (b *commentRepository) getFindAllCommentQuery() string {
	return `select c.id, c.review_id, c.content, u.id, u.name
from comments as c
left join users as u
on c.create_user_id = u.id
where c.delete_flag = 0 %s
LIMIT 20 OFFSET ?`
}

func (b *commentRepository) FindAllComment(page int, search string) (entity.Comments, error) {
	result := make(entity.Comments, 0)
	query := b.getFindAllCommentQuery()
	if search == "" {
		query = fmt.Sprintf(query, "")
	} else {
		query = fmt.Sprintf(query, "AND name like '%"+search+"%'")
	}

	rows, err := b.handler.Query(query, (page-1)*20)
	if err != nil {
		return nil, &model.BsError{
			Code: model.EINTERNAL,
			Op:   "b.handler.Query",
			Err:  err,
		}
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
			&comment.UserID,
			&comment.UserName,
		); err != nil {
			return nil, &model.BsError{
				Code: model.EINTERNAL,
				Op:   "row.Scan",
				Err:  err,
			}
		}
		result = append(result, &comment)
	}
	return result, nil
}

func (b *commentRepository) getFindCommentByIDQuery() string {
	return `select c.id, c.review_id, c.content, u.id, u.name
from comments as c
left join users as u
on c.create_user_id = u.id
where c.id = ? c.delete_flag = 0`
}

func (b *commentRepository) FindCommentByID(id int) (*entity.Comment, error) {
	query := b.getFindCommentByIDQuery()
	row := b.handler.QueryRow(query, id)
	var comment entity.Comment
	if err := row.Scan(
		&comment.ID,
		&comment.ReviewID,
		&comment.Content,
		&comment.UserID,
		&comment.UserName,
	); err != nil {
		return nil, &model.BsError{
			Code: model.EINTERNAL,
			Op:   "row.Scan",
			Err:  err,
		}
	}
	return &comment, nil
}

func (b *commentRepository) getCreateCommentQuery() string {
	return `INSERT INTO comments (review_id, content, create_user_id) VALUES (?, ?, ?);`
}

func (b *commentRepository) CreateComment(userID int, comment entity.Comment) (*entity.Comment, error) {
	query := b.getCreateCommentQuery()
	result, err := b.handler.Exec(query,
		comment.ReviewID,
		comment.Content,
		userID,
	)
	if err != nil {
		return nil, &model.BsError{
			Code: model.EINTERNAL,
			Op:   "b.handler.Exec",
			Err:  err,
		}
	}
	insID, err := result.LastInsertId()
	if err != nil {
		return nil, &model.BsError{
			Code: model.EINTERNAL,
			Op:   "result.LastInsertId",
			Err:  err,
		}
	}
	comment.ID = int(insID)
	return &comment, nil
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
		return &model.BsError{
			Code: model.EINTERNAL,
			Op:   "b.handler.Exec",
			Err:  err,
		}
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
		return &model.BsError{
			Code: model.EINTERNAL,
			Op:   "b.handler.Exec",
			Err:  err,
		}
	}
	return nil
}

func (b *commentRepository) getFindCommentByReviewID() string {
	return `select c.id, c.review_id, c.content, u.id, u.name
from comments as c
left join users as u
on c.create_user_id = u.id
where c.delete_flag = 0 and c.review_id = ?`
}

func (b *commentRepository) FindCommentByReviewID(reviewID int) (entity.Comments, error) {
	result := make(entity.Comments, 0)
	query := b.getFindCommentByReviewID()
	rows, err := b.handler.Query(query, reviewID)
	if err != nil {
		return nil, &model.BsError{
			Code: model.EINTERNAL,
			Op:   "b.handler.Query",
			Err:  err,
		}
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
			&comment.UserID,
			&comment.UserName,
		); err != nil {
			return nil, &model.BsError{
				Code: model.EINTERNAL,
				Op:   "rows.Scan",
				Err:  err,
			}
		}
		result = append(result, &comment)
	}
	return result, nil
}
