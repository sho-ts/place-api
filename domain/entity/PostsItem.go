package entity

import (
	"time"
)

type PostsItem struct {
	PostId    string    `json:"postId"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"createdAt"`
	Thumbnail string    `json:"thumbnail"`
	User      User      `json:"user"`
}

func NewPostsItem(
	postId string,
	caption string,
	createdAt time.Time,
	thumbnail string,
	user User,
) PostsItem {
	return PostsItem{
		PostId:    postId,
		Caption:   caption,
		CreatedAt: createdAt,
		Thumbnail: thumbnail,
		User:      user,
	}
}
