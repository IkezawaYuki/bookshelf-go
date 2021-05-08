package entity

import "time"

type Reviews []*Review

type Review struct {
	ID          int
	BookID      int
	Title       string
	Content     string
	ReadingDate time.Time
	Model
}
