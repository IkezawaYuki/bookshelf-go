package entity

type Comments []*Comment

type Comment struct {
	ID       int
	ReviewID int
	Content  string
}
