package controller

import (
	"fmt"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/outputport"
	"net/http"
	"strconv"
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

func (ctr *BookshelfController) GetBooks(c outputport.Context) error {
	panic("implement")
}

func (ctr *BookshelfController) GetBook(c outputport.Context) error {
	bookID := c.QueryParam("id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("param=%s, %v", bookID, err))
		return err
	}

	book, err := ctr.bookInputport.FindBookByID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	if book == nil {
		_ = c.JSON(http.StatusNotFound, fmt.Errorf("id :%d", id))
		return err
	}

	return c.JSON(http.StatusOK, book)
}
