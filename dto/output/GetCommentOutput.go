package output

import (
	"github.com/sho-ts/place-api/entity"
	"time"
)

type GetCommentOutput struct {
	CommentId string      `json:"commentId"`
	PostId    string      `json:"postId"`
	Content   string      `json:"content"`
	CreatedAt time.Time   `json:"createdAt"`
	User      entity.User `json:"user"`
}

func NewGetCommentOutput(
	commentId string,
	postId string,
	content string,
	createdAt time.Time,
	user entity.User,
) GetCommentOutput {
	return GetCommentOutput{
		CommentId: commentId,
		PostId:    postId,
		Content:   content,
		CreatedAt: createdAt,
		User:      user,
	}
}
