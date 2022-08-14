package output

import (
	"github.com/sho-ts/place-api/entity"
	"time"
)

type GetCommentOutput struct {
	Id        string      `json:"commentId"`
	PostId    string      `json:"postId"`
	Content   string      `json:"content"`
	CreatedAt time.Time   `json:"createdAt"`
	User      entity.User `json:"user"`
}
