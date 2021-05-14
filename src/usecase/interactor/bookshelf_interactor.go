package interactor

import (
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
