package adapter

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
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
	return ``
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
	return ``
}

func (b *commentRepository) FindCommentByID(id int) (comment entity.Comment, err error) {
	query := b.getFindCommentByIDQuery()
	row := b.handler.QueryRow(query, id)
	if err = row.Scan(
		&comment.ID,
		&comment.ReviewID,
		&comment.Content,
	); err != nil {
		return
	}
	return
}

func (b *commentRepository) getCreateCommentQuery() string {
	return ``
}

func (b *commentRepository) CreateComment(userID int, comment entity.Comment) error {
	query := b.getCreateCommentQuery()
	_, err := b.handler.Exec(query,
		comment.ID,
		comment.ReviewID,
		comment.Content,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (b *commentRepository) getUpdateCommentQuery() string {
	return ``
}

func (b *commentRepository) UpdateComment(userID int, comment entity.Comment) error {
	query := b.getUpdateCommentQuery()
	_, err := b.handler.Exec(query,
		comment.ID,
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
	return ``
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
