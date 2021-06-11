package outputport

import "time"

type OutputPort interface {
	ShowBook(c Context) error
}

type BookDetail struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Publisher   string    `json:"publisher"`
	Author      string    `json:"author"`
	DateOfIssue time.Time `json:"date_of_issue"`
	Price       float64   `json:"price"`
	Reviews     []Review  `json:"reviews"`
}

type Review struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	User        string    `json:"user"`
	Content     string    `json:"content"`
	ReadingDate time.Time `json:"reading_date"`
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
