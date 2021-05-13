package adapter

import (
	"github.com/IkezawaYuki/bookshelf-go/src/domain/entity"
	"github.com/IkezawaYuki/bookshelf-go/src/domain/repository"
	"github.com/IkezawaYuki/bookshelf-go/src/interfaces/datastore"
)

func NewUserRepository(handler datastore.DBHandler) repository.UserRepository {
	return &userRepository{handler: handler}
}

type userRepository struct {
	handler datastore.DBHandler
}

func (r *userRepository) getFindAllUserQuery() string {
	return `select name, gender, birthday, email, occupation_code, address_code
from users where delete_flag = 0`
}

func (r *userRepository) FindAllUser() (entity.Users, error) {
	result := make(entity.Users, 0)
	query := r.getFindAllUserQuery()
	rows, err := r.handler.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()
	for rows.Next() {
		user := entity.User{}
		if err := rows.Scan(
			&user.Name,
			&user.Gender,
			&user.BirthDate,
			&user.Email,
			&user.OccupationCode,
			&user.AddressCode,
		); err != nil {
			return nil, err
		}
		result = append(result, &user)
	}
	return result, nil
}

func (r *userRepository) getFindUserByIDQuery() string {
	return `select name, gender, birthday, email, occupation_code, address_code
from users where id = ? and delete_flag = 0`
}

func (r *userRepository) FindUserByID(id int) (user entity.User, err error) {
	query := r.getFindUserByIDQuery()
	row := r.handler.QueryRow(query, id)
	err = row.Scan(
		&user.Name,
		&user.Gender,
		&user.BirthDate,
		&user.Email,
		&user.OccupationCode,
		&user.AddressCode,
	)
	return
}

func (r *userRepository) getCreateUserQuery() string {
	return `INSERT INTO users (name,
gender,
birthday,
email,
occupation_code,
address_code,
create_user_id)
VALUES
(?, ?, ?, ?, ?, ?, ?);`
}

func (r *userRepository) CreateUser(userID int, user entity.User) error {
	query := r.getCreateUserQuery()
	_, err := r.handler.Exec(query,
		user.Name,
		user.Gender,
		user.BirthDate,
		user.Email,
		user.OccupationCode,
		user.AddressCode,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) getUpdateUserQuery() string {
	return `UPDATE users SET
name = ?,
gender = ?,
birthday = ?,
email = ?,
occupation_code = ?,
address_code = ?,
update_user_id = ?`
}

func (r *userRepository) UpdateUser(userID int, user entity.User) error {
	query := r.getUpdateUserQuery()
	_, err := r.handler.Exec(query,
		user.Name,
		user.Gender,
		user.BirthDate,
		user.Email,
		user.OccupationCode,
		user.AddressCode,
		userID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) getDeleteUserByIDQuery() string {
	return `UPDATE users
SET delete_user_id = ?,
delete_date = now(),
delete_flag = 1
WHERE id = ?`
}

func (r *userRepository) DeleteUserByID(userID int, id int) error {
	query := r.getDeleteUserByIDQuery()
	_, err := r.handler.Exec(query,
		userID,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}
