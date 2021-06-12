package entity

import (
	"time"
)

type Books []*Book

type Book struct {
	ID          int
	Name        string
	Publisher   string
	Author      string
	DateOfIssue time.Time
	Price       float64
	Model
}

func (b *Book) TaxIncludedPrice() int {
	return int(b.Price * tax)
}

func (b *Book) GetTitle() string {
	return b.Name
}

func (b *Book) GetBody() string {
	// todo string builder を使う
	return b.Author
}
