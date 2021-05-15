package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

func NewBookshelfInteractor(
	bookRepo repository.BookRepository,
	commentRepo repository.CommentRepository,
	reviewRepo repository.ReviewRepository,
	shelfRepo repository.ShelfRepository,
	userRepo repository.UserRepository,
) inputport.BookShelfInputPort {
	return &bookshelfInteractor{
		bookRepo:    bookRepo,
		commentRepo: commentRepo,
		reviewRepo:  reviewRepo,
		shelfRepo:   shelfRepo,
		userRepo:    userRepo,
	}
}

type bookshelfInteractor struct {
	bookRepo    repository.BookRepository
	commentRepo repository.CommentRepository
	reviewRepo  repository.ReviewRepository
	shelfRepo   repository.ShelfRepository
	userRepo    repository.UserRepository
}

func (b *bookshelfInteractor) FindAllBook() (entity.Books, error) {
	return b.bookRepo.FindAllBook()
}

func (b *bookshelfInteractor) FindBookByID(id int) (entity.Book, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) CreateBook(userID int, book entity.Book) error {
	panic("implement me")
}

func (b *bookshelfInteractor) UpdateBook(userID int, book entity.Book) error {
	panic("implement me")
}

func (b *bookshelfInteractor) DeleteBookByID(userID int, id int) error {
	panic("implement me")
}

func (b *bookshelfInteractor) FindAllComment() (entity.Comments, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) FindCommentByID(id int) (entity.Comment, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) CreateComment(userID int, book entity.Comment) error {
	panic("implement me")
}

func (b *bookshelfInteractor) UpdateComment(userID int, book entity.Comment) error {
	panic("implement me")
}

func (b *bookshelfInteractor) DeleteCommentByID(userID int, id int) error {
	panic("implement me")
}

func (b *bookshelfInteractor) FindAllReview() (entity.Reviews, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) FindReviewByID(id int) (entity.Review, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) CreateReview(userID int, review entity.Review) error {
	panic("implement me")
}

func (b *bookshelfInteractor) UpdateReview(userID int, review entity.Review) error {
	panic("implement me")
}

func (b *bookshelfInteractor) DeleteReviewByID(userID int, id int) error {
	panic("implement me")
}

func (b *bookshelfInteractor) FindAllShelf() (entity.Shelves, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) FindShelfByID(id int) (entity.Shelf, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) CreateShelf(userID int, shelf entity.Shelf) error {
	panic("implement me")
}

func (b *bookshelfInteractor) UpdateShelf(userID int, shelf entity.Shelf) error {
	panic("implement me")
}

func (b *bookshelfInteractor) DeleteShelfByID(userID int, id int) error {
	panic("implement me")
}

func (b *bookshelfInteractor) FindAllUser() (entity.Users, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) FindUserByID(id int) (entity.User, error) {
	panic("implement me")
}

func (b *bookshelfInteractor) CreateUser(userID int, user entity.User) error {
	panic("implement me")
}

func (b *bookshelfInteractor) UpdateUser(userID int, user entity.User) error {
	panic("implement me")
}

func (b *bookshelfInteractor) DeleteUserByID(userID int, id int) error {
	panic("implement me")
}
