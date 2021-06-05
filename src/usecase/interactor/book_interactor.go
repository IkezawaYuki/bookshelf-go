package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type bookInteractor struct {
	bookRepo repository.BookRepository
}

func NewBookInteractor(repo repository.BookRepository) inputport.BookInputPort {
	return &bookInteractor{
		bookRepo: repo,
	}
}

func (b *bookInteractor) FindAllBook() (entity.Books, error) {
	return b.bookRepo.FindAllBook()
}

func (b *bookInteractor) FindBookByID(id int) (*entity.Book, error) {
	return b.bookRepo.FindBookByID(id)
}

func (b *bookInteractor) CreateBook(userID int, book entity.Book) (*entity.Book, error) {
	return b.bookRepo.CreateBook(userID, book)
}

func (b *bookInteractor) UpdateBook(userID int, book entity.Book) error {
	return b.bookRepo.UpdateBook(userID, book)
}

func (b *bookInteractor) DeleteBookByID(userID int, id int) error {
	return b.bookRepo.DeleteBookByID(userID, id)
}
