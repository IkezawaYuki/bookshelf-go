package entity

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/model"
	"time"
)

type Books []*Book

type Book struct {
	ID          int
	Name        model.Name
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
	return string(b.Name)
}
