package outputport

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"time"
)

type Presenter interface {
	ConvertBook(book *entity.Book, reviews entity.Reviews) *BookDetail
}

func NewPresenter() Presenter {
	return &presenter{}
}

type presenter struct {
}

func (p *presenter) ConvertBook(book *entity.Book, reviews entity.Reviews) *BookDetail {
	var detail BookDetail
	detail.ID = book.ID
	detail.Name = book.Name
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
