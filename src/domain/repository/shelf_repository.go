package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type ShelfRepository interface {
	FindAllShelf() (entity.Shelves, error)
	FindShelfByID(id int) (entity.Shelf, error)
	CreateShelf(book entity.Shelf) error
	UpdateShelf(book entity.Shelf) error
	DeleteShelfByID(id int) error
}
