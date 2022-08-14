package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/entity"
)

func CreateComment(i input.CreateCommentInput) (entity.Comment, error) {
	comment := entity.Comment{
		Id:      i.Id,
		UserId:  i.UserId,
		PostId:  i.PostId,
		Content: i.Content,
	}

	result := database.DB.Create(&comment)

	return comment, result.Error
}