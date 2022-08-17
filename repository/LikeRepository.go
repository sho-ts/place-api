package repository

import (
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/infrastructure/database/table"
	"github.com/sho-ts/place-api/util"
)

type LikeRepository struct{}

func NewLikeRepository() LikeRepository {
	return LikeRepository{}
}

func (repository LikeRepository) Store(postId string, userId string) error {
	like := table.Like{
		Id:     util.GetULID(),
		PostId: postId,
		UserId: userId,
	}

	result := database.DB.Create(&like)

	return result.Error
}

func (repository LikeRepository) Remove(postId string, userId string) error {
	result := database.DB.
		Where("post_id = ?", postId).
		Where("user_id = ?", userId).
		Delete(&table.Like{})

	return result.Error
}

func (repository LikeRepository) CheckDuplicate(postId string, userId string) (
	bool, error,
) {
	var count int64

	result := database.DB.
		Where("post_id = ?", postId).
		Where("user_id = ?", userId).
		Count(&count)

	return count > 0, result.Error
}
