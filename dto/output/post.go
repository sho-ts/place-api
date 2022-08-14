package output

import (
	"github.com/sho-ts/place-api/entity"
)

type GetPostOutput struct {
	PostId  string           `json:"postId"`
	UserId  string           `json:"userId"`
	Caption string           `json:"caption"`
	Files   []entity.Storage `json:"files"`
}

type GetPostsOutput struct {
	PostId    string `json:"postId"`
	UserId    string `json:"userId"`
	Caption   string `json:"caption"`
	Thumbnail string `json:"thumbnail"`
}
