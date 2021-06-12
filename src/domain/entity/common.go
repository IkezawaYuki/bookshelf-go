package entity

import "time"

const tax = 0.1

type Model struct {
	CreateUserID int
	CreateDate   time.Time
	UpdateUserID *int
	UpdateDate   time.Time
	DeleteUserID *int
	DeleteDate   *time.Time
	DeleteFlag   int
}
