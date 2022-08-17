package entity

import (
	"time"
)

type Post struct {
	PostId         string          `json:"postId"`
	Caption        string          `json:"caption"`
	CreatedAt      time.Time       `json:"createdAt"`
	Liked          int             `json:"liked"`
	StorageObjects []StorageObject `json:"files"`
	User           User            `json:"user"`
}

func NewPost(
	postId string,
	caption string,
	liked int,
	storageObjects []StorageObject,
	user User,
) Post {
	return Post{
		PostId:         postId,
		Caption:        caption,
		Liked:          liked,
		StorageObjects: storageObjects,
		User:           user,
	}
}
