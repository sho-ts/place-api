package repository

import (
	"github.com/sho-ts/place-api/application/util"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/infrastructure/database/table"
	"strings"
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

func (repository FollowRepository) GetFollowsByDisplayId(
	displayId string, limit int, offset int,
) ([]entity.User, int64, error) {
	var items []entity.User
	var count int64

	queryBase := database.DB.
		Select(strings.Join([]string{
			"users.id as Id",
			"users.display_id as DisplayId",
			"users.name as Name",
			"users.avatar as Avatar",
		}, ",")).
		Table("follows").
		Joins("join users on users.id = follows.follow_user_id").
		Where("follows.follower_user_id = (select id from users where display_id = ? limit 1)", displayId)

	result := queryBase.Limit(limit).Offset(offset).Scan(&items)

	if result.Error != nil {
		return items, count, result.Error
	}

	result = queryBase.Count(&count)

	return items, count, result.Error
}

func (repository FollowRepository) GetFollowersByDisplayId(
	displayId string, limit int, offset int,
) ([]entity.User, int64, error) {
	var items []entity.User
	var count int64

	queryBase := database.DB.
		Select(strings.Join([]string{
			"users.id as Id",
			"users.display_id as DisplayId",
			"users.name as Name",
			"users.avatar as Avatar",
		}, ",")).
		Table("follows").
		Joins("join users on users.id = follows.follower_user_id").
		Where("follows.follow_user_id = (select id from users where display_id = ? limit 1)", displayId)

	result := queryBase.Limit(limit).Offset(offset).Scan(&items)

	if result.Error != nil {
		return items, count, result.Error
	}

	result = queryBase.Count(&count)

	return items, count, result.Error
}
