package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type shelfInteractor struct {
}

func NewShelfInteractor() inputport.ShelfInputPort {
	return &shelfInteractor{}
}

func (s *shelfInteractor) FindAllShelf() (entity.Shelves, error) {
	panic("implement me")
}

func (s *shelfInteractor) FindShelfByID(id int) (entity.Shelf, error) {
	panic("implement me")
}

func (s *shelfInteractor) CreateShelf(userID int, shelf entity.Shelf) (entity.Shelf, error) {
	panic("implement me")
}

func (s *shelfInteractor) UpdateShelf(userID int, shelf entity.Shelf) error {
	panic("implement me")
}

func (s *shelfInteractor) DeleteShelfByID(userID int, id int) error {
	panic("implement me")
}
