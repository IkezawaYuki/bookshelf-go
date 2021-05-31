package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/adapter"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type bookInteractor struct {
}

func NewBookInteractor() inputport.BookInputPort {
	return &bookInteractor{}
}

func (b *bookInteractor) FindAllBook() (entity.Books, error) {
	return adapter.BookRepo.FindAllBook()
}

func (b *bookInteractor) FindBookByID(id int) (*entity.Book, error) {
	return adapter.BookRepo.FindBookByID(id)
}

func (b *bookInteractor) CreateBook(userID int, book entity.Book) (*entity.Book, error) {
	return adapter.BookRepo.CreateBook(userID, book)
}

func (b *bookInteractor) UpdateBook(userID int, book entity.Book) error {
	return adapter.BookRepo.UpdateBook(userID, book)
}

func (b *bookInteractor) DeleteBookByID(userID int, id int) error {
	return adapter.BookRepo.DeleteBookByID(userID, id)
}
