package outputport

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"time"
)

type Presenter interface {
	ConvertBookDetail(book *entity.Book, reviews entity.Reviews) *BookDetail
	ConvertUser(user *entity.User) *User
	ConvertUsers(user entity.Users) Users
	ConvertReview(review *entity.Review) *Review
	ConvertReviewDetail(review *entity.Review, comments entity.Comments) *ReviewDetail
	ConvertComment(comment *entity.Comment) *Comment
}

func NewPresenter() Presenter {
	return &presenter{}
}

type presenter struct {
}

func (p *presenter) ConvertBookDetail(book *entity.Book, reviews entity.Reviews) *BookDetail {
	var detail BookDetail
	detail.ID = book.ID
	detail.Name = string(book.Name)
	detail.Publisher = book.Publisher
	detail.Author = book.Author
	detail.DateOfIssue = book.DateOfIssue.Format("2006-01-02")

	for _, r := range reviews {
		var rev Review
		rev.ID = r.ID
		rev.Title = r.Title
		rev.Content = r.Content
		rev.ReadingDate = r.ReadingDate.Format("2006-01-02")
		detail.Reviews = append(detail.Reviews, rev)
	}
	return &detail
}

func (p *presenter) ConvertUser(user *entity.User) *User {
	var u User
	u.ID = user.ID
	u.Name = user.Name
	u.Gender = user.GetGender()
	u.BirthDate = user.BirthDate.Format("2006-01-02")
	u.Email = user.Email
	u.Occupation = user.OccupationName
	u.Address = user.AddressName
	return &u
}

func (p *presenter) ConvertUsers(users entity.Users) Users {
	var result Users
	for _, u := range users {
		result = append(result, p.ConvertUser(u))
	}
	return result
}

func (p *presenter) ConvertReview(review *entity.Review) *Review {
	return &Review{
		ID:          review.ID,
		Title:       review.Title,
		User:        "todo",
		Content:     review.Content,
		ReadingDate: review.ReadingDate.Format("2006-01-02"),
	}
}

func (p *presenter) ConvertReviewDetail(review *entity.Review, comments entity.Comments) *ReviewDetail {
	panic("implement me")
}

func (p *presenter) ConvertComment(comment *entity.Comment) *Comment {
	return &Comment{
		ID:       comment.ID,
		ReviewID: comment.ReviewID,
		UserID:   comment.UserID,
		Content:  comment.Content,
	}
}

type BookDetail struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Publisher   string   `json:"publisher"`
	Author      string   `json:"author"`
	DateOfIssue string   `json:"date_of_issue"`
	Price       float64  `json:"price"`
	Reviews     []Review `json:"reviews"`
}

type Review struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	User        string `json:"user"`
	Content     string `json:"content"`
	ReadingDate string `json:"reading_date"`
}

type Books []BookIndex

type BookIndex struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Publisher   string    `json:"publisher"`
	Author      string    `json:"author"`
	DateOfIssue time.Time `json:"date_of_issue"`
	Price       float64   `json:"price"`
}

type Users []*User

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
	Email      string `json:"email"`
	Occupation string `json:"occupation,omitempty"`
	Address    string `json:"address,omitempty"`
}

type ReviewDetail struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	User        string    `json:"user"`
	Content     string    `json:"content"`
	ReadingDate string    `json:"reading_date"`
	Comments    []Comment `json:"comments"`
}

type Comment struct {
	ID       int    `json:"id"`
	ReviewID int    `json:"review_id"`
	UserID   int    `json:"user_id"`
	Content  string `json:"content"`
}
