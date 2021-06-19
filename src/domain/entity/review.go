package entity

import "time"

type Reviews []*Review

type Review struct {
	ID          int
	BookID      int
	UserID      int
	UserName    string
	Title       string
	Content     string
	ReadingDate time.Time
	Model
}
