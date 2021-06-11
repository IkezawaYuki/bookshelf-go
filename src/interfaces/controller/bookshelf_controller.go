package controller

import (
	"fmt"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/aggregate"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/logger"
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
	presenter        outputport.OutputPort
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
	return c.JSON(http.StatusOK, "0.0.0")
}

func (ctr *BookshelfController) GetBooks(c outputport.Context) error {
	return c.JSON(http.StatusOK, "book")
}

// GetBook 本の取得
// @Title GetBook
// @Description idによる本の取得
// @Accept json
// @Produce json
// @Param id path int true "本のID"
// @Success 200 {object} entity.Book
// @Router /book/{id} [get]
func (ctr *BookshelfController) GetBook(c outputport.Context) error {
	logger.Info("GetBook is invoked")
	bookID := c.Param("id")
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

func (ctr *BookshelfController) RegisterBook(c outputport.Context) error {
	var book entity.Book
	err := c.Bind(book)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	userID := c.Get("user_id").(int)
	insBook, err := ctr.bookInputport.CreateBook(userID, book)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusCreated, insBook)
}

func (ctr *BookshelfController) UpdateBook(c outputport.Context) error {
	var book entity.Book
	err := c.Bind(book)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	userID := c.Get("user_id").(int)
	if err := ctr.bookInputport.UpdateBook(userID, book); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusAccepted, nil)
}

func (ctr *BookshelfController) DeleteBook(c outputport.Context) error {
	bookID := c.QueryParam("id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("param=%s, %v", bookID, err))
		return err
	}

	userID := c.Get("user_id").(int)
	if err := ctr.bookInputport.DeleteBookByID(userID, id); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusAccepted, nil)
}

func (ctr *BookshelfController) FindReviews(c outputport.Context) error {
	bookID := c.QueryParam("book_id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("param=%s, %v", bookID, err))
		return err
	}

	reviews, err := ctr.reviewInputport.FindReviewByBookID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusOK, reviews)
}

func (ctr *BookshelfController) BookIndex(c outputport.Context) error {
	return nil
}

func (ctr *BookshelfController) BookShow(c outputport.Context) error {
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

	reviews, err := ctr.reviewInputport.FindReviewByBookID(book.ID)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	var bookDf aggregate.BookDf
	bookDf.Book = book
	bookDf.Reviews = reviews

	return c.JSON(http.StatusOK, bookDf)
}
