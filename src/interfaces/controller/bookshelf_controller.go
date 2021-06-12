package controller

import (
	"fmt"
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
	presenter        outputport.Presenter
}

func NewBookshelfController(
	bookInputport inputport.BookInputPort,
	commentInputport inputport.CommentInputPort,
	reviewInputport inputport.ReviewInputPort,
	shelfInputport inputport.ShelfInputPort,
	userInputport inputport.UserInputPort,
	presenter outputport.Presenter,
) BookshelfController {
	return BookshelfController{
		bookInputport:    bookInputport,
		commentInputport: commentInputport,
		reviewInputport:  reviewInputport,
		shelfInputport:   shelfInputport,
		userInputport:    userInputport,
		presenter:        presenter,
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
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("param=%s, %v", bookID, err).Error())
		return err
	}

	book, err := ctr.bookInputport.FindBookByID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	if book == nil {
		_ = c.JSON(http.StatusNotFound, fmt.Errorf("id :%d", id).Error())
		return err
	}

	return c.JSON(http.StatusOK, book)
}

// RegisterBook 本の登録
// @Title RegisterBook
// @Description 本の登録
// @Accept json
// @Produce json
// @Success 200 {object} entity.Book
// @Router /book [post]
func (ctr *BookshelfController) RegisterBook(c outputport.Context) error {
	var book entity.Book
	err := c.Bind(book)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	userID := c.Get("user_id").(int)
	insBook, err := ctr.bookInputport.CreateBook(userID, book)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusCreated, insBook)
}

func (ctr *BookshelfController) UpdateBook(c outputport.Context) error {
	var book entity.Book
	err := c.Bind(book)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	userID := c.Get("user_id").(int)
	if err := ctr.bookInputport.UpdateBook(userID, book); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusAccepted, nil)
}

func (ctr *BookshelfController) DeleteBook(c outputport.Context) error {
	bookID := c.QueryParam("id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("param=%s, %v", bookID, err).Error())
		return err
	}

	userID := c.Get("user_id").(int)
	if err := ctr.bookInputport.DeleteBookByID(userID, id); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusAccepted, nil)
}

func (ctr *BookshelfController) FindReviews(c outputport.Context) error {
	bookID := c.QueryParam("book_id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("param=%s, %v", bookID, err).Error())
		return err
	}

	reviews, err := ctr.reviewInputport.FindReviewByBookID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, reviews)
}

func (ctr *BookshelfController) BookIndex(c outputport.Context) error {
	return nil
}

func (ctr *BookshelfController) ShowBook(c outputport.Context) error {
	logger.Info("ShowBook is invoked")

	bookID := c.Param("id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("param=%s, %v", bookID, err).Error())
		return err
	}

	book, err := ctr.bookInputport.FindBookByID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	reviews, err := ctr.reviewInputport.FindReviewByBookID(book.ID)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, ctr.presenter.ConvertBook(book, reviews))
}

func (ctr *BookshelfController) ShowUser(c outputport.Context) error {
	logger.Info("ShowUser is invoked")

	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("param=%s, %v", userID, err).Error())
		return err
	}

	user, err := ctr.userInputport.FindUserByID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, ctr.presenter.ConvertUser(user))
}

func (ctr *BookshelfController) GetUsers(c outputport.Context) error {
	logger.Info("GetUsers is invoked")

	users, err := ctr.userInputport.FindAllUser()
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ctr.presenter.ConvertUsers(users))
}

func (ctr *BookshelfController) OutputUsersReport(c outputport.Context) error {
	panic("implement me")
}
