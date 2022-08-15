package object

import (
	"errors"
	"unicode/utf8"
)

type UserName struct {
	Value string
}

func NewUserName(value string) UserName {
	return UserName{
		Value: value,
	}
}

func (di UserName) Valid() error {
	len := utf8.RuneCountInString(di.Value)

	if len > 32 {
		return errors.New("ユーザー名は32文字以内にしてください")
	}

	return nil
}