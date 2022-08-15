package object

import (
	"errors"
	"unicode/utf8"
)

type CommentContent struct {
	Value string
}

func NewCommentContent(value string) CommentContent {
	return CommentContent{
		Value: value,
	}
}

func (cc CommentContent) Valid() error {
	len := utf8.RuneCountInString(cc.Value)

	if len > 500 {
		return errors.New("文字数は500文字以内にしてください")
	}

	return nil
}
