package entity

type Comments []*Comment

type Comment struct {
	ID       int
	ReviewID int
	UserID   int
	Content  string
	Model
}
