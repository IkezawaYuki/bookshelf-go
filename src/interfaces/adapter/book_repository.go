package adapter

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
	"strings"
)

func NewBookRepository(handler datastore.DBHandler) repository.BookRepository {
	return &bookRepository{handler: handler}
}

type bookRepository struct {
	handler datastore.DBHandler
}

func (b *bookRepository) getFindAllBookQuery() string {
	return `SELECT 
id,
name,
publisher,
author,
date_of_issue,
price
FROM books WHERE delete_flag = 0`
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
		book := entity.Book{}
		if err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Publisher,
			&book.Author,
			&book.DateOfIssue,
			&book.Price,
		); err != nil {
			return nil, err
		}
		result = append(result, &book)
	}
	return result, nil
}

func (b *bookRepository) getFindBookByIDQuery() string {
	return `SELECT 
id,
name,
publisher,
author,
date_of_issue,
price
FROM books WHERE id = ? AND delete_flag = 0`
}

func (b *bookRepository) FindBookByID(id int) (*entity.Book, error) {
	query := b.getFindBookByIDQuery()
	row := b.handler.QueryRow(query, id)
	var book entity.Book
	if err := row.Scan(
		&book.ID,
		&book.Name,
		&book.Publisher,
		&book.Author,
		&book.DateOfIssue,
		&book.Price,
	); err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *bookRepository) getCreateBookQuery() string {
	return `INSERT INTO books(name, publisher, author, date_of_issue, price, create_user_id) 
VALUES (?, ?, ?, ?, ?, ?)`
}

func (b *bookRepository) CreateBook(userID int, book entity.Book) (*entity.Book, error) {
	query := b.getCreateBookQuery()
	result, err := b.handler.Exec(query,
		book.Name,
		book.Publisher,
		book.Author,
		book.DateOfIssue,
		book.Price,
		userID,
	)
	if err != nil {
		return nil, err
	}
	insID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	book.ID = int(insID)
	return &book, nil
}

func (b *bookRepository) getUpdateBookQuery() string {
	return `UPDATE books
SET name = ?, 
publisher = ?, 
author = ?, 
date_of_issue = ?, 
price = ?, 
update_user_id = ?, 
WHERE id = ?`
}

func (b *bookRepository) UpdateBook(userID int, book entity.Book) error {
	query := b.getUpdateBookQuery()
	_, err := b.handler.Exec(query,
		book.Name,
		book.Publisher,
		book.Author,
		book.DateOfIssue,
		book.Price,
		userID,
		book.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (b *bookRepository) getDeleteBookByIDQuery() string {
	return `UPDATE books
SET delete_user_id = ?,
delete_date = now(),
delete_flag = 1
WHERE id = ?`
}

func (b *bookRepository) DeleteBookByID(userID int, id int) error {
	query := b.getDeleteBookByIDQuery()
	_, err := b.handler.Exec(query,
		userID,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (b *bookRepository) getFindByNameQuery() string {
	return `SELECT 
id,
name,
publisher,
author,
date_of_issue,
price
FROM books WHERE delete_flag = 0
and name like ("%$word%")`
}

func (b *bookRepository) FindByName(name string) (entity.Books, error) {
	result := make(entity.Books, 0)
	query := strings.ReplaceAll(b.getFindByNameQuery(), "$word", name)
	rows, err := b.handler.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		book := entity.Book{}
		if err := rows.Scan(
			&book.ID,
			&book.Name,
			&book.Publisher,
			&book.Author,
			&book.DateOfIssue,
			&book.Price,
		); err != nil {
			return nil, err
		}
		result = append(result, &book)
	}
	return result, nil
}
