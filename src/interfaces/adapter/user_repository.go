package adapter

import (
	"database/sql"
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
	return `select u.id, u.name, u.gender, u.birthday, u.email, u.occupation_code, o.name, u.address_code, a.name
from users as u 
left join occupation as o
on u.occupation_code = o.code and o.delete_flag = 0
left join address as a
on u.address_code = a.code and a.delete_flag = 0
where u.delete_flag = 0`
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
		var occupationCode sql.NullString
		var occupationName sql.NullString
		var addressCode sql.NullString
		var addressName sql.NullString
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Gender,
			&user.BirthDate,
			&user.Email,
			&occupationCode,
			&occupationName,
			&addressCode,
			&addressName,
		); err != nil {
			return nil, err
		}
		user.OccupationCode = occupationCode.String
		user.OccupationName = occupationName.String
		user.AddressCode = addressCode.String
		user.AddressName = addressName.String
		result = append(result, &user)
	}
	return result, nil
}

func (r *userRepository) getFindUserByIDQuery() string {
	return `select u.name, u.gender, u.birthday, u.email, u.occupation_code, o.name, u.address_code, a.name
from users as u 
left join occupation as o
on u.occupation_code = o.code and o.delete_flag = 0
left join address as a
on u.address_code = a.code and a.delete_flag = 0
where u.id = ? and u.delete_flag = 0
order by u.id`
}

func (r *userRepository) FindUserByID(id int) (*entity.User, error) {
	query := r.getFindUserByIDQuery()
	row := r.handler.QueryRow(query, id)
	var user entity.User
	var occupationCode sql.NullString
	var occupationName sql.NullString
	var addressCode sql.NullString
	var addressName sql.NullString
	err := row.Scan(
		&user.Name,
		&user.Gender,
		&user.BirthDate,
		&user.Email,
		&occupationCode,
		&occupationName,
		&addressCode,
		&addressName,
	)
	user.ID = id
	user.OccupationCode = occupationCode.String
	user.OccupationName = occupationName.String
	user.AddressCode = addressCode.String
	user.AddressName = addressName.String
	return &user, err
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

func (r *userRepository) CreateUser(userID int, user entity.User) (*entity.User, error) {
	query := r.getCreateUserQuery()
	result, err := r.handler.Exec(query,
		user.Name,
		user.Gender,
		user.BirthDate,
		user.Email,
		user.OccupationCode,
		user.AddressCode,
		userID,
	)
	if err != nil {
		return nil, err
	}
	insID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.ID = int(insID)
	return &user, nil
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

func (r *userRepository) getFindUserByEmailQuery() string {
	return `select id, name, gender, birthday, occupation_code, address_code
from users where email = ? and delete_flag = 0`
}

func (r *userRepository) FindUserByEmail(email string) (*entity.User, error) {
	query := r.getFindUserByEmailQuery()
	row := r.handler.QueryRow(query, email)
	var user entity.User
	var occupationCode sql.NullString
	var addressCode sql.NullString
	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Gender,
		&user.BirthDate,
		&occupationCode,
		&addressCode,
	)
	user.OccupationCode = occupationCode.String
	user.AddressCode = addressCode.String
	return &user, err
}
