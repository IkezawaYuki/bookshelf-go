package entity

import "time"

type Reviews []*Review

type Review struct {
	ID          int
	ReviewerID  int
	Title       string
	Content     string
	ReadingDate time.Time
}
