package entity

import "time"

type BookReviews []*BookReview

type BookReview struct {
	ID          int
	ReviewerID  int
	Title       string
	Content     string
	ReadingDate time.Time
}
