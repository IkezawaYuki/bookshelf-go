package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type ShelfRepository interface {
	FindAllShelf() (entity.Shelves, error)
	FindShelfByID(id int) (entity.Shelf, error)
	CreateShelf(userID int, shelf entity.Shelf) (entity.Shelf, error)
	UpdateShelf(userID int, shelf entity.Shelf) error
	DeleteShelfByID(userID int, id int) error
}
