package entity

type Shelves []*Shelf

type Shelf struct {
	ID      int
	Name    string
	OwnerID int
	Model
}
