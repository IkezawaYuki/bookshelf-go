package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type shelfInteractor struct {
	shelfRepo repository.ShelfRepository
}

func NewShelfInteractor(repo repository.ShelfRepository) inputport.ShelfInputPort {
	return &shelfInteractor{
		shelfRepo: repo,
	}
}

func (s *shelfInteractor) FindAllShelf() (entity.Shelves, error) {
	return s.shelfRepo.FindAllShelf()
}

func (s *shelfInteractor) FindShelfByID(id int) (*entity.Shelf, error) {
	return s.shelfRepo.FindShelfByID(id)
}

func (s *shelfInteractor) CreateShelf(userID int, shelf entity.Shelf) (*entity.Shelf, error) {
	return s.shelfRepo.CreateShelf(userID, shelf)
}

func (s *shelfInteractor) UpdateShelf(userID int, shelf entity.Shelf) error {
	return s.shelfRepo.UpdateShelf(userID, shelf)
}

func (s *shelfInteractor) DeleteShelfByID(userID int, id int) error {
	return s.shelfRepo.DeleteShelfByID(userID, id)
}
