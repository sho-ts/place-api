package object

import (
	"errors"
	"unicode/utf8"
)

type DisplayId struct {
	Value string
}

func NewDisplayId(value string) DisplayId {
	return DisplayId{
		Value: value,
	}
}

func (di DisplayId) Valid() error {
	len := utf8.RuneCountInString(di.Value)

	if len > 24 {
		return errors.New("IDは24文字以内にしてください")
	}

	return nil
}
