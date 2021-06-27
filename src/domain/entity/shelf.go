package entity

import "github.com/IkezawaYuki/bookshelf-go/src/domain/model"

type Shelves []*Shelf

type Shelf struct {
	ID      int
	OwnerID int
	Name    model.Name
	Model
}
