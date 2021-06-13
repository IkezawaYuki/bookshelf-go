package entity

import (
	"fmt"
	"time"
)

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

func (u Users) Header() []interface{} {
	return []interface{}{
		"ID",
		"氏名",
		"性別",
		"生年月日",
		"メールアドレス",
		"職業コード",
		"職業名称",
		"都道府県コード",
		"都道府県名称",
		"更新日",
	}
}

func (u Users) Cells() [][]interface{} {
	var cells [][]interface{}
	cells = append(cells, u.Header())
	for _, user := range u {
		row := make([]interface{}, 0)
		row = append(row, fmt.Sprintf("%d", user.ID))
		row = append(row, user.Name)
		row = append(row, user.GetGender())
		row = append(row, user.BirthDate.Format("2006-01-02"))
		row = append(row, user.Email)
		row = append(row, user.OccupationCode)
		row = append(row, user.OccupationName)
		row = append(row, user.AddressCode)
		row = append(row, user.AddressName)
		row = append(row, user.UpdateDate.Format("2006-01-02 15:04:05"))
		cells = append(cells, row)
	}
	return cells
}

func (u Users) SheetName() string {
	return fmt.Sprintf("ユーザー情報_%s", CurrentTimeJST().Format("20060102"))
}
