package inputport

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type BookShelfInputPort interface {
	FindAllBook() (entity.Books, error)
	FindBookByID(id int) (entity.Book, error)
	CreateBook(userID int, book entity.Book) (entity.Book, error)
	UpdateBook(userID int, book entity.Book) error
	DeleteBookByID(userID int, id int) error
	FindAllComment() (entity.Comments, error)
	FindCommentByID(id int) (entity.Comment, error)
	CreateComment(userID int, book entity.Comment) (entity.Comment, error)
	UpdateComment(userID int, book entity.Comment) error
	DeleteCommentByID(userID int, id int) error
	FindAllReview() (entity.Reviews, error)
	FindReviewByID(id int) (entity.Review, error)
	CreateReview(userID int, review entity.Review) error
	UpdateReview(userID int, review entity.Review) error
	DeleteReviewByID(userID int, id int) error
	FindAllShelf() (entity.Shelves, error)
	FindShelfByID(id int) (entity.Shelf, error)
	CreateShelf(userID int, shelf entity.Shelf) error
	UpdateShelf(userID int, shelf entity.Shelf) error
	DeleteShelfByID(userID int, id int) error
	FindAllUser() (entity.Users, error)
	FindUserByID(id int) (entity.User, error)
	CreateUser(userID int, user entity.User) error
	UpdateUser(userID int, user entity.User) error
	DeleteUserByID(userID int, id int) error
}
