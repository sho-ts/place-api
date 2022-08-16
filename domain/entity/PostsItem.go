package entity


import (
	"github.com/sho-ts/place-api/entity"
	"time"
)

type PostsItem struct {
	PostId    string      `json:"postId"`
	Caption   string      `json:"caption"`
	CreatedAt time.Time   `json:"createdAt"`
	Thumbnail string      `json:"thumbnail"`
	User      entity.User `json:"user"`
}

func NewPostsItem(
	postId string,
	caption string,
	createdAt time.Time,
	thumbnail string,
	user entity.User,
) PostsItem {
	return PostsItem{
		PostId:    postId,
		Caption:   caption,
		CreatedAt: createdAt,
		Thumbnail: thumbnail,
		User:      user,
	}
}