package output

import (
	"github.com/sho-ts/place-api/entity"
	"time"
)

type GetPostOutput struct {
	PostId    string           `json:"postId"`
	Caption   string           `json:"caption"`
	CreatedAt time.Time        `json:"createdAt"`
	Liked     int              `json:"liked"`
	Files     []entity.Storage `json:"files"`
	User      entity.User      `json:"user"`
}

type GetPostsOutput struct {
	PostId    string      `json:"postId"`
	Caption   string      `json:"caption"`
	CreatedAt time.Time   `json:"createdAt"`
	Thumbnail string      `json:"thumbnail"`
	User      entity.User `json:"user"`
}
