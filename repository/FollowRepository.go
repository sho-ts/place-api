package repository

import (
	"github.com/sho-ts/place-api/application/util"
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/infrastructure/database/table"
)

type FollowRepository struct{}

func NewFollowRepository() FollowRepository {
	return FollowRepository{}
}

func (repository FollowRepository) Store(followUserId string, followerUserId string) error {
	follow := table.Follow{
		Id:             util.GetULID(),
		FollowUserId:   followUserId,
		FollowerUserId: followerUserId,
	}

	result := database.DB.Create(&follow)

	return result.Error
}

func (repository FollowRepository) Remove(followUserId string, followerUserId string) error {
	result := database.DB.
		Where("follow_user_id = ?", followUserId).
		Where("follower_user_id = ?", followerUserId).
		Delete(&table.Follow{})

	return result.Error
}

func (repository FollowRepository) CheckDuplicate(followUserId string, followerUserId string) (bool, error) {
	var count int64

	result := database.DB.
		Table("follows").
		Where("follow_user_id = ?", followUserId).
		Where("follower_user_id = ?", followerUserId).
		Count(&count)

	return count > 0, result.Error
}
