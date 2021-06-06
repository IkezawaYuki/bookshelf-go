package interactor

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/usecase/inputport"
)

type userInteractor struct {
	userRepo repository.UserRepository
}

func NewUserInteractor(repo repository.UserRepository) inputport.UserInputPort {
	return &userInteractor{
		userRepo: repo,
	}
}

func (u *userInteractor) FindAllUser() (entity.Users, error) {
	return u.userRepo.FindAllUser()
}

func (u *userInteractor) FindUserByID(id int) (*entity.User, error) {
	return u.userRepo.FindUserByID(id)
}

func (u *userInteractor) CreateUser(userID int, user entity.User) (*entity.User, error) {
	return u.userRepo.CreateUser(userID, user)
}

func (u *userInteractor) UpdateUser(userID int, user entity.User) error {
	return u.userRepo.UpdateUser(userID, user)
}

func (u *userInteractor) DeleteUserByID(userID int, id int) error {
	return u.userRepo.DeleteUserByID(userID, id)
}

func (u *userInteractor) FindUserByEmail(email string) (*entity.User, error) {
	return u.userRepo.FindUserByEmail(email)
}
