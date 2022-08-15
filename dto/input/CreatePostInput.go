package input

import (
	"github.com/sho-ts/place-api/object"
)

type CreatePostInput struct {
	PostId  string
	UserId  string
	Caption object.PostCaption
	Urls    []string
}

func NewCreatePostInput(
	postId string,
	userId string,
	caption string,
	urls []string,
) (CreatePostInput, error) {
	postCaption := object.NewPostCaption(caption)
	err := postCaption.Valid()

	createPostInput := CreatePostInput{
		PostId:  postId,
		UserId:  userId,
		Caption: postCaption,
		Urls:    urls,
	}
	return createPostInput, err
}
