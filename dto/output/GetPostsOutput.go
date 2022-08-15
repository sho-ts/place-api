package output

import (
	"github.com/sho-ts/place-api/entity"
	"time"
)

type GetPostsOutputItem struct {
	PostId    string      `json:"postId"`
	Caption   string      `json:"caption"`
	CreatedAt time.Time   `json:"createdAt"`
	Thumbnail string      `json:"thumbnail"`
	User      entity.User `json:"user"`
}

func NewGetPostsOutputItem(
	postId string,
	caption string,
	createdAt time.Time,
	thumbnail string,
	user entity.User,
) GetPostsOutputItem {
	return GetPostsOutputItem{
		PostId:    postId,
		Caption:   caption,
		CreatedAt: createdAt,
		Thumbnail: thumbnail,
		User:      user,
	}
}

type GetPostsOutput struct {
	Items     []GetPostsOutputItem `json:"items"`
}

func NewGetPostsOutput(items []GetPostsOutputItem) GetPostsOutput {
	return GetPostsOutput{
		Items: items,
	}
}
