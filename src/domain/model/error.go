package model

import (
	"bytes"
	"fmt"
)

const (
	EINTERNAL = "internal"
)

type BsError struct {
	Code    string
	Message string
	Op      string
	Err     error
}

func ErrorCode(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*BsError); e.Code != "" {
		return e.Code
	} else if ok && e.Err != nil {
		return ErrorCode(e.Err)
	}
	return EINTERNAL
}

func ErrorMessage(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*BsError); e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return ErrorMessage(e.Err)
	}
	return EINTERNAL
}

func (e *BsError) Error() string {
	var buf bytes.Buffer
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		if e.Code != "" {
			fmt.Fprintf(&buf, "<%s> ", e.Code)
		}
		buf.WriteString(e.Message)
	}
	return buf.String()
}
