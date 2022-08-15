package object

import (
	"errors"
	"unicode/utf8"
)

type PostCaption struct {
	Value string
}

func NewPostCaption(value string) PostCaption {
	return PostCaption{
		Value: value,
	}
}

func (pc PostCaption) Valid() error {
	len := utf8.RuneCountInString(pc.Value)

	if len > 500 {
		return errors.New("文字数は500文字以内にしてください")
	}

	return nil
}
