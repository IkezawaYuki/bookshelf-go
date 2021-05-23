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
	return b.bookRepo.FindBookByID(id)
}

func (b *bookshelfInteractor) CreateBook(userID int, book entity.Book) (entity.Book, error) {
	return b.bookRepo.CreateBook(userID, book)
}

func (b *bookshelfInteractor) UpdateBook(userID int, book entity.Book) error {
	return b.bookRepo.UpdateBook(userID, book)
}

func (b *bookshelfInteractor) DeleteBookByID(userID int, id int) error {
	return b.bookRepo.DeleteBookByID(userID, id)
}

func (b *bookshelfInteractor) FindAllComment() (entity.Comments, error) {
	return b.commentRepo.FindAllComment()
}

func (b *bookshelfInteractor) FindCommentByID(id int) (entity.Comment, error) {
	return b.commentRepo.FindCommentByID(id)
}

func (b *bookshelfInteractor) CreateComment(userID int, book entity.Comment) (entity.Comment, error) {
	return b.commentRepo.CreateComment(userID, book)
}

func (b *bookshelfInteractor) UpdateComment(userID int, book entity.Comment) error {
	return b.commentRepo.UpdateComment(userID, book)
}

func (b *bookshelfInteractor) DeleteCommentByID(userID int, id int) error {
	return b.commentRepo.DeleteCommentByID(userID, id)
}

func (b *bookshelfInteractor) FindAllReview() (entity.Reviews, error) {
	return b.reviewRepo.FindAllReview()
}

func (b *bookshelfInteractor) FindReviewByID(id int) (entity.Review, error) {
	return b.reviewRepo.FindReviewByID(id)
}

func (b *bookshelfInteractor) CreateReview(userID int, review entity.Review) (entity.Review, error) {
	return b.reviewRepo.CreateReview(userID, review)
}

func (b *bookshelfInteractor) UpdateReview(userID int, review entity.Review) error {
	return b.reviewRepo.UpdateReview(userID, review)
}

func (b *bookshelfInteractor) DeleteReviewByID(userID int, id int) error {
	return b.reviewRepo.DeleteReviewByID(userID, id)
}

func (b *bookshelfInteractor) FindAllShelf() (entity.Shelves, error) {
	return b.shelfRepo.FindAllShelf()
}

func (b *bookshelfInteractor) FindShelfByID(id int) (entity.Shelf, error) {
	return b.shelfRepo.FindShelfByID(id)
}

func (b *bookshelfInteractor) CreateShelf(userID int, shelf entity.Shelf) (entity.Shelf, error) {
	return b.shelfRepo.CreateShelf(userID, shelf)
}

func (b *bookshelfInteractor) UpdateShelf(userID int, shelf entity.Shelf) error {
	return b.shelfRepo.UpdateShelf(userID, shelf)
}

func (b *bookshelfInteractor) DeleteShelfByID(userID int, id int) error {
	return b.shelfRepo.DeleteShelfByID(userID, id)
}

func (b *bookshelfInteractor) FindAllUser() (entity.Users, error) {
	return b.userRepo.FindAllUser()
}

func (b *bookshelfInteractor) FindUserByID(id int) (entity.User, error) {
	return b.userRepo.FindUserByID(id)
}

func (b *bookshelfInteractor) CreateUser(userID int, user entity.User) (entity.User, error) {
	return b.userRepo.CreateUser(userID, user)
}

func (b *bookshelfInteractor) UpdateUser(userID int, user entity.User) error {
	return b.userRepo.UpdateUser(userID, user)
}

func (b *bookshelfInteractor) DeleteUserByID(userID int, id int) error {
	return b.userRepo.DeleteUserByID(userID, id)
}
