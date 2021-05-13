package entity

type Shelves []*Shelf

type Shelf struct {
	ID      int
	OwnerID int
	Name    string
	Model
}
