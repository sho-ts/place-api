package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/entity"
)

func CreatePost(postId string, authId string, caption string) (entity.Post, error) {
	post := entity.Post{
		Id:      postId,
		UserId:  authId,
		Caption: caption,
	}

	result := database.DB.Create(&post)

	return post, result.Error
}
