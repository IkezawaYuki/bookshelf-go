package entity

import "time"

type Users []*User

type User struct {
	ID             int
	Name           string
	Gender         int
	BirthDate      time.Time
	Email          string
	OccupationCode string
	OccupationName string
	AddressCode    string
	AddressName    string
	Model
}

func (u *User) GetGender() string {
	if u.Gender == 0 {
		return "男性"
	} else if u.Gender == 1 {
		return "女性"
	}
	return ""
}

func (u Users) Header() []string {
	return []string{
		"ID",
		"氏名",
		"性別",
		"生年月日",
		"メールアドレス",
		"職業コード",
		"職業名称",
		"都道府県コード",
		"都道府県名称",
	}
}

func (u Users) Cells() [][]string {
	panic("implement me")
}

func (u Users) SheetName() string {
	panic("implement me")
}

func (u Users) FileName() string {
	panic("implement me")
}
