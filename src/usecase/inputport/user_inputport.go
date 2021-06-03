package inputport

import "github.com/IkezawaYuki/bookshelf-go/src/domain/entity"

type UserInputPort interface {
	FindAllUser() (entity.Users, error)
	FindUserByID(id int) (*entity.User, error)
	CreateUser(userID int, user entity.User) (*entity.User, error)
	UpdateUser(userID int, user entity.User) error
	DeleteUserByID(userID int, id int) error
	FindUserByEmail(email string) (*entity.User, error)
}
