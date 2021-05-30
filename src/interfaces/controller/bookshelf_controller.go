package controller

import (
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"
)

type BookshelfController struct {
	bookInputport    inputport.BookInputPort
	commentInputport inputport.CommentInputPort
	reviewInputport  inputport.ReviewInputPort
	shelfInputport   inputport.ShelfInputPort
	userInputport    inputport.UserInputPort
}

func NewBookshelfController(
	bookInputport inputport.BookInputPort,
	commentInputport inputport.CommentInputPort,
	reviewInputport inputport.ReviewInputPort,
	shelfInputport inputport.ShelfInputPort,
	userInputport inputport.UserInputPort,
) BookshelfController {
	return BookshelfController{
		bookInputport:    bookInputport,
		commentInputport: commentInputport,
		reviewInputport:  reviewInputport,
		shelfInputport:   shelfInputport,
		userInputport:    userInputport,
	}
}

func (ctr *BookshelfController) GetVersion(c outputport.Context) error {
	return nil
}
