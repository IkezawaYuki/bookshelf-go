package adapter

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
)

func NewRepository(handler datastore.DBHandler) repository.BookRepository {
	return &bookRepository{handler: handler}
}

type bookRepository struct {
	handler datastore.DBHandler
}

func (b *bookRepository) getFindAllBookQuery() string {
	return ""
}

func (b *bookRepository) FindAllBook() (entity.Books, error) {
	result := make(entity.Books, 0)
	query := b.getFindAllBookQuery()
	rows, err := b.handler.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		book := new(entity.Book)
		if err := rows.Scan(); err != nil {
			return nil, err
		}
		result = append(result, book)
	}
	return result, nil
}

func (b *bookRepository) FindBookByID(id int) (entity.Book, error) {
	panic("implement me")
}

func (b *bookRepository) CreateBook(book entity.Book) error {
	panic("implement me")
}

func (b *bookRepository) UpdateBook(book entity.Book) error {
	panic("implement me")
}

func (b *bookRepository) DeleteBookByID(id int) error {
	panic("implement me")
}
