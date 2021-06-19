package entity

import "time"

type Comments []*Comment

type Comment struct {
	ID          int
	ReviewID    int
	UserID      int
	Content     string
	CommentDate time.Time
	Model
}
