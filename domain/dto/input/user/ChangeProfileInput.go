package input

import "mime/multipart"

type ChangeProfileInput struct {
	UserId    string
	DisplayId string
	Name      string
	File      multipart.File
	FileName  string
}

func NewChangeProfileInput(
	userId string,
	displayId string,
	userName string,
	file multipart.File,
	fileName string,
) ChangeProfileInput {
	changeProfileInput := ChangeProfileInput{
		UserId:    userId,
		DisplayId: displayId,
		Name:      userName,
		File:      file,
		FileName:  fileName,
	}

	return changeProfileInput
}
