package entity

import "time"

type Users []*User

type User struct {
	ID         int
	Name       string
	Gender     int
	BirthDate  time.Time
	Email      string
	Occupation string
	Address    string
}
