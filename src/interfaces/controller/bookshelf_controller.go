package controller

import (
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"
	"net/http"
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

// GetVersion バージョン情報の取得
// @Summary バージョン情報を文字列で返す
// @Success 200 {string} 0.0.0
// @Router /version [get]
func (ctr *BookshelfController) GetVersion(c outputport.Context) error {
	return c.JSON(http.StatusOK, "0,0,0")
}

func (ctr *BookshelfController) GetBook(c outputport.Context) error {
	panic("implement")
}
