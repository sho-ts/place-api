package entity

import (
	"time"
)

type Comment struct {
	CommentId string    `json:"commentId"`
	PostId    string    `json:"postId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user"`
}

func NewComment(
	commentId string,
	postId string,
	content string,
	user User,
) Comment {
	return Comment{
		CommentId: commentId,
		PostId:    postId,
		Content:   content,
		User:      user,
	}
}
