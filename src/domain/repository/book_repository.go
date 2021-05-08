package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type BookRepository interface {
	FindAllBook() (entity.Books, error)
	FindBookByID(id int) (entity.Book, error)
	CreateBook(book entity.Book) error
	UpdateBook(book entity.Book) error
	DeleteBookByID(id int) error
}
