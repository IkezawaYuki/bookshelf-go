package entity

type Comments []*Comment

type Comment struct {
	ID       int
	ReviewID int
	UserID   int
	UserName string
	Content  string
	Model
}
