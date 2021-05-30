package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/adapter"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type userInteractor struct {
}

func NewUserInteractor() inputport.UserInputPort {
	return &userInteractor{}
}

func (u *userInteractor) FindAllUser() (entity.Users, error) {
	panic("implement me")
}

func (u *userInteractor) FindUserByID(id int) (entity.User, error) {
	panic("implement me")
}

func (u *userInteractor) CreateUser(userID int, user entity.User) (*entity.User, error) {
	panic("implement me")
}

func (u *userInteractor) UpdateUser(userID int, user entity.User) error {
	panic("implement me")
}

func (u *userInteractor) DeleteUserByID(userID int, id int) error {
	panic("implement me")
}

func (u *userInteractor) FindUserByEmail(email string) (*entity.User, error) {
	return adapter.UserRepo.FindUserByEmail(email)
}
