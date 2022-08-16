package input

import "mime/multipart"

type CreatePostInput struct {
	PostId   string
	UserId   string
	Caption  string
	File     multipart.File
	FileName string
}

func NewCreatePostInput(
	postId string,
	userId string,
	caption string,
	file multipart.File,
	fileName string,
) CreatePostInput {
	return CreatePostInput{
		PostId:   postId,
		UserId:   userId,
		Caption:  caption,
		File:     file,
		FileName: fileName,
	}
}
