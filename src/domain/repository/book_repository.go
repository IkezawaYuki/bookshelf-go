package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type BookRepository interface {
	FindAllBook() (entity.Books, error)
	FindBookByID(id int) (entity.Book, error)
	CreateBook(userID int, book entity.Book) (entity.Book, error)
	UpdateBook(userID int, book entity.Book) error
	DeleteBookByID(userID int, id int) error
	FindByName(name string) (entity.Books, error)
}
