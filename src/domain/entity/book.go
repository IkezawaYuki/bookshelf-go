package entity

import "time"

type Books []*Book

type Book struct {
	ID          int
	Name        string
	Publisher   string
	Author      string
	DateOfIssue time.Time
	Price       float64
}

func (b *Book) TaxIncludedPrice() int {
	return int(b.Price * tax)
}
