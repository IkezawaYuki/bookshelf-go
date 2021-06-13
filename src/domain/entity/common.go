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

func CurrentTimeJST() *time.Time {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	current := time.Now().In(loc)
	return &current
}
