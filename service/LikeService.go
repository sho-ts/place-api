package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
)

type LikeService struct {}

func NewLikeService () LikeService {
  likeService := LikeService{}
  return likeService
}

func (ls LikeService) AddLike(i input.HandleLikeInput) error {
	like := entity.Like{
		Id:     util.GetULID(),
		PostId: i.PostId,
		UserId: i.UserId,
	}

	result := database.DB.Create(&like)

	return result.Error
}

func (ls LikeService) RemoveLike(i input.HandleLikeInput) error {
	var like entity.Like

	result := database.DB.
		Where("post_id = ?", i.PostId).
		Where("user_id = ?", i.UserId).
		Delete(&like)

	return result.Error
}

func (ls LikeService) GetLikeCount(postId string) (output.CountOutput, error) {
	var count int64

	result := database.DB.
		Table("likes").
		Where("post_id = ?", postId).
		Count(&count)

	o := output.CountOutput{
		Count: count,
	}

	return o, result.Error
}

func (ls LikeService) CheckDuplicateLike(i input.HandleLikeInput) (bool, error) {
	var count int64

	result := database.DB.
		Table("likes").
		Where("post_id = ?", i.PostId).
		Where("user_id = ?", i.UserId).
		Count(&count)

	return count > 0, result.Error
}