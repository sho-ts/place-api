package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/output"
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

func GetPost(postId string) (output.GetPostResponseOutput, error) {
	var post entity.Post
	result := database.DB.Where("id = ?", postId).First(&post)

	var files []entity.Storage

	result = database.DB.Where("post_id = ?", postId).Find(&files)

	o := output.GetPostResponseOutput{
		PostId:  post.Id,
		UserId:  post.UserId,
		Caption: post.Caption,
		Files:   files,
	}

	return o, result.Error
}
