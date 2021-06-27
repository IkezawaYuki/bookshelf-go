package model

import "strings"

type Name string

func NewName(str string) Name {
	if len(strings.Split("", str)) > 50 {
		str = strings.Join(strings.Split("", str)[:50], "")
	}
	return Name(str)
}

type Content string

func NewContent(str string) Content {
	if len(strings.Split("", str)) > 250 {
		str = strings.Join(strings.Split("", str)[:250], "")
	}
	return Content(str)
}
