package entity

import "time"

const tax = 0.1

type Model struct {
	CreateUserID int        `json:"create_user_id"`
	CreateDate   time.Time  `json:"create_date"`
	UpdateUserID *int       `json:"update_user_id,omitempty"`
	UpdateDate   time.Time  `json:"update_date"`
	DeleteUserID *int       `json:"delete_user_id,omitempty"`
	DeleteDate   *time.Time `json:"delete_date,omitempty"`
	DeleteFlag   int        `json:"delete_flag"`
}
