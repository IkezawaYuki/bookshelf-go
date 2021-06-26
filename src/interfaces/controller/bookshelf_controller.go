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
	bookInputport     inputport.BookInputPort
	commentInputport  inputport.CommentInputPort
	reviewInputport   inputport.ReviewInputPort
	shelfInputport    inputport.ShelfInputPort
	userInputport     inputport.UserInputPort
	presenter         outputport.Presenter
	spreadsheetClient outputport.SpreadsheetOutputPort
}

func NewBookshelfController(
	bookInputport inputport.BookInputPort,
	commentInputport inputport.CommentInputPort,
	reviewInputport inputport.ReviewInputPort,
	shelfInputport inputport.ShelfInputPort,
	userInputport inputport.UserInputPort,
	presenter outputport.Presenter,
	spreadsheetClient outputport.SpreadsheetOutputPort,
) BookshelfController {
	return BookshelfController{
		bookInputport:    bookInputport,
		commentInputport: commentInputport,
		reviewInputport:  reviewInputport,
		shelfInputport:   shelfInputport,
		userInputport:    userInputport,
		presenter:        presenter,

		spreadsheetClient: spreadsheetClient,
	}
}

func (ctr *BookshelfController) HealthCheck(c outputport.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}

// GetVersion バージョン情報の取得
// @Summary バージョン情報を文字列で返す
// @Success 200 {string} 0.0.0
// @Router /version [get]
func (ctr *BookshelfController) GetVersion(c outputport.Context) error {
	return c.JSON(http.StatusOK, "0.0.0")
}

func (ctr *BookshelfController) GetBooks(c outputport.Context) error {
	queryPage := c.QueryParam("page")
	page := 1
	if queryPage != "" {
		page, _ = strconv.Atoi(queryPage)
	}
	name := c.QueryParam("search")

	books, err := ctr.bookInputport.FindAllBook(page, name)
	if err != nil {
		return err
	}
	// todo
	fmt.Println(books)
	fmt.Println("todo")
	return c.JSON(http.StatusOK, "books")
}

// GetBook 本の取得
// @Title GetBook
// @Summary idによる本の取得
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

// CreateBook 本の登録
// @Title CreateBook
// @Summary 本の登録
// @Description 本の登録
// @Accept json
// @Produce json
// @Success 200 {object} entity.Book
// @Router /book [post]
func (ctr *BookshelfController) CreateBook(c outputport.Context) error {
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

// UpdateBook 本の登録
// @Title UpdateBook
// @Summary 本の更新
// @Description 本の更新
// @Accept json
// @Produce json
// @Success 200 {object} entity.Book
// @Router /book [post]
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

// DeleteBook 本の登録
// @Title DeleteBook
// @Summary 本の削除
// @Description 本の削除
// @Accept json
// @Param id path int true "id"
// @Success 202
// @Router /book/{id} [delete]
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

// ShowBook 本の詳細情報の取得
// @Title ShowBook
// @Summary 本の詳細情報の取得
// @Description 本の詳細情報の取得
// @Accept json
// @Param id path int true "id"
// @Success 200 {object} outputport.BookDetail
// @Router /book/detail/{id} [get]
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

	return c.JSON(http.StatusOK, ctr.presenter.ConvertBookDetail(book, reviews))
}

// ShowUser ユーザー情報の取得
// @Title ShowUser
// @Summary ユーザー情報の取得
// @Description ユーザー情報の取得
// @Accept json
// @Param id path int true "id"
// @Success 200 {object} outputport.User
// @Router /user/detail/{id} [get]
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

// GetUsers ユーザー情報の全員取得
// @Title GetUsers
// @Summary ユーザー情報の全員取得
// @Description ユーザー情報の全員取得
// @Accept json
// @Success 200 {object} outputport.Users
// @Router /users [get]
func (ctr *BookshelfController) GetUsers(c outputport.Context) error {
	logger.Info("GetUsers is invoked")

	users, err := ctr.userInputport.FindAllUser()
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, ctr.presenter.ConvertUsers(users))
}

// OutputUsersReport ユーザー情報全員のスプレッドシート出力
// @Title OutputUsersReport
// @Summary ユーザー情報全員のスプレッドシート出力
// @Description ユーザー情報全員のスプレッドシート出力
// @Accept json
// @Success 200 {string} url
// @Router /users/report [get]
func (ctr *BookshelfController) OutputUsersReport(c outputport.Context) error {
	logger.Info("OutputUsersReport is invoked")

	refreshToken := c.Get("refresh_token").(string)

	users, err := ctr.userInputport.FindAllUser()
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	url, err := ctr.spreadsheetClient.OutputOneSheet(refreshToken, "filename", users)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusOK, url)
}

// GetReview レビューの取得
// @Title GetReview
// @Summary idによるレビューの取得
// @Description idによるレビューの取得
// @Accept json
// @Produce json
// @Param id path int true "レビューのID"
// @Success 200 {object} outputport.Review
// @Router /review/{id} [get]
func (ctr *BookshelfController) GetReview(c outputport.Context) error {
	reviewID := c.Param("id")
	id, err := strconv.Atoi(reviewID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Sprintf("err: %v, reviewID=%s", err, reviewID))
		return err
	}

	review, err := ctr.reviewInputport.FindReviewByID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusOK, ctr.presenter.ConvertReview(review))
}

func (ctr *BookshelfController) ShowReview(c outputport.Context) error {
	reviewID := c.Param("id")
	id, err := strconv.Atoi(reviewID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	review, err := ctr.reviewInputport.FindReviewByID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	comments, err := ctr.commentInputport.FindCommentByReviewID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusOK, ctr.presenter.ConvertReviewDetail(review, comments))
}

func (ctr *BookshelfController) GetReviews(c outputport.Context) error {
	panic("implement me")
}

func (ctr *BookshelfController) UpdateReview(c outputport.Context) error {
	var review entity.Review
	err := c.Bind(review)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	userID := c.Get("user_id").(int)

	if err := ctr.reviewInputport.UpdateReview(userID, review); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusOK, review)
}

func (ctr *BookshelfController) CreateReview(c outputport.Context) error {
	var review entity.Review
	err := c.Bind(review)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	userID := c.Get("user_id").(int)

	insReview, err := ctr.reviewInputport.CreateReview(userID, review)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusCreated, insReview)
}

// DeleteReview レビューの登録
// @Title DeleteReview
// @Summary レビューの削除
// @Description レビューの削除
// @Accept json
// @Param id path int true "id"
// @Success 202
// @Router /review/{id} [delete]
func (ctr *BookshelfController) DeleteReview(c outputport.Context) error {
	reviewID := c.QueryParam("id")
	id, err := strconv.Atoi(reviewID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, fmt.Errorf("reviewID=%s, err=%v", reviewID, err).Error())
		return err
	}

	userID := c.Get("user_id").(int)
	if err := ctr.reviewInputport.DeleteReviewByID(userID, id); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusAccepted, nil)
}

// GetComment コメントの取得
// @Title GetComment
// @Summary idによるコメントの取得
// @Description idによるコメントの取得
// @Accept json
// @Produce json
// @Param id path int true "コメントのID"
// @Success 200 {object} outputport.Comment
// @Router /comment/{id} [get]
func (ctr *BookshelfController) GetComment(c outputport.Context) error {
	commentID := c.Param("id")
	id, err := strconv.Atoi(commentID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	comment, err := ctr.commentInputport.FindCommentByID(id)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusAccepted, ctr.presenter.ConvertComment(comment))
}

func (ctr *BookshelfController) GetComments(c outputport.Context) error {
	panic("implement me")
}

func (ctr *BookshelfController) CreateComment(c outputport.Context) error {
	var comment entity.Comment
	if err := c.Bind(comment); err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	userID := c.Get("user_id").(int)
	insComment, err := ctr.commentInputport.CreateComment(userID, comment)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusOK, insComment)
}

func (ctr *BookshelfController) UpdateComment(c outputport.Context) error {
	var comment entity.Comment
	if err := c.Bind(comment); err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	userID := c.Get("user_id").(int)
	if err := ctr.commentInputport.UpdateComment(userID, comment); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusOK, nil)
}

// DeleteComment コメントの登録
// @Title DeleteComment
// @Summary コメントの削除
// @Description コメントの削除
// @Accept json
// @Param id path int true "id"
// @Success 202
// @Router /comment/{id} [delete]
func (ctr *BookshelfController) DeleteComment(c outputport.Context) error {
	commentID := c.Param("id")
	id, err := strconv.Atoi(commentID)
	if err != nil {
		_ = c.JSON(http.StatusBadRequest, err)
		return err
	}

	userID := c.Get("user_id").(int)
	if err := ctr.commentInputport.DeleteCommentByID(userID, id); err != nil {
		_ = c.JSON(http.StatusInternalServerError, err)
		return err
	}

	return c.JSON(http.StatusAccepted, nil)
}
