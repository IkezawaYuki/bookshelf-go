package repository

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type UserRepository interface {
	FindAllUser() (entity.Users, error)
	FindUserByID(id int) (entity.User, error)
	CreateUser(book entity.User) error
	UpdateUser(book entity.User) error
	DeleteUserByID(id int) error
}
