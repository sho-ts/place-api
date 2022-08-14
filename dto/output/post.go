package output

import (
	"github.com/sho-ts/place-api/entity"
)

type GetPostOutput struct {
	PostId  string           `json:"postId"`
	Caption string           `json:"caption"`
	Files   []entity.Storage `json:"files"`
	User    entity.User      `json:"user"`
}

type GetPostsOutput struct {
	PostId    string      `json:"postId"`
	Caption   string      `json:"caption"`
	Thumbnail string      `json:"thumbnail"`
	User      entity.User `json:"user"`
}
