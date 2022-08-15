package input

import (
	"github.com/sho-ts/place-api/object"
)

type CreateCommentInput struct {
	Id      string
	UserId  string
	PostId  string
	Content object.CommentContent
}

func NewCreateCommentInput(
	id string,
	userId string,
	postId string,
	content string,
) (CreateCommentInput, error) {
	commentContent := object.NewCommentContent(content)
	err := commentContent.Valid()

	createCommentInput := CreateCommentInput{
		Id:      id,
		UserId:  userId,
		PostId:  postId,
		Content: commentContent,
	}

	return createCommentInput, err
}
