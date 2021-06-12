package adapter

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
)

func NewShelfRepository(handler datastore.DBHandler) repository.ShelfRepository {
	return &shelfRepository{handler: handler}
}

type shelfRepository struct {
	handler datastore.DBHandler
}

func (r *shelfRepository) getFindAllShelfQuery() string {
	return `select id, owner_id, name from shelves where delete_flag = 0`
}

func (r *shelfRepository) FindAllShelf() (entity.Shelves, error) {
	result := make(entity.Shelves, 0)
	query := r.getFindAllShelfQuery()
	rows, err := r.handler.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		shelf := entity.Shelf{}
		if err := rows.Scan(
			&shelf.ID,
			&shelf.OwnerID,
			&shelf.Name,
		); err != nil {
			return nil, err
		}
		result = append(result, &shelf)
	}
	return result, nil
}

func (r *shelfRepository) getFindShelfByIDQuery() string {
	return `select id, owner_id, name from shelves where id = ? and delete_flag = 0`
}

func (r *shelfRepository) FindShelfByID(id int) (*entity.Shelf, error) {
	query := r.getFindShelfByIDQuery()
	row := r.handler.QueryRow(query, id)
	var shelf entity.Shelf
	err := row.Scan(
		&shelf.ID,
		&shelf.OwnerID,
		&shelf.Name,
	)
	return &shelf, err
}

func (r *shelfRepository) getCreateShelfQuery() string {
	return `INSERT INTO shelves (owner_id, name, create_user_id) VALUES (?, ?, ?);`
}

func (r *shelfRepository) CreateShelf(userID int, shelf entity.Shelf) (*entity.Shelf, error) {
	query := r.getCreateShelfQuery()
	result, err := r.handler.Exec(query,
		shelf.OwnerID,
		shelf.Name,
		userID,
	)
	if err != nil {
		return nil, err
	}
	insID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	shelf.ID = int(insID)
	return &shelf, nil
}

func (r *shelfRepository) getUpdateShelfQuery() string {
	return `UPDATE shelves SET
owner_id = ?,
name = ?,
create_user_id = ?
WHERE id = ?`
}

func (r *shelfRepository) UpdateShelf(userID int, shelf entity.Shelf) error {
	query := r.getUpdateShelfQuery()
	_, err := r.handler.Exec(query,
		shelf.OwnerID,
		shelf.Name,
		userID,
		shelf.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *shelfRepository) getDeleteShelfByIDQuery() string {
	return `UPDATE shelves
SET delete_user_id = ?,
delete_date = now(),
delete_flag = 1
WHERE id = ?`
}

func (r *shelfRepository) DeleteShelfByID(userID int, id int) error {
	query := r.getDeleteShelfByIDQuery()
	_, err := r.handler.Exec(query,
		userID,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
