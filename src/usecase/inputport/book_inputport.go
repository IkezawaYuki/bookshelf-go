package inputport

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type BookInputPort interface {
	FindAllBook(page int, search string) (entity.Books, error)
	FindBookByID(id int) (*entity.Book, error)
	CreateBook(userID int, book entity.Book) (*entity.Book, error)
	UpdateBook(userID int, book entity.Book) error
	DeleteBookByID(userID int, id int) error
}
