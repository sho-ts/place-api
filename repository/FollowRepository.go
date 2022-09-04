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
	displayId string, userId string, limit int, offset int,
) ([]entity.Follow, int64, error) {
	var items []entity.Follow
	var count int64

	s := strings.Join([]string{
		"users.id AS Id",
		"users.display_id AS DisplayId",
		"users.name AS Name",
		"users.avatar AS Avatar",
	}, ",")

	if userId != "" {
		s = s + ", CASE WHEN users.follow_user_id IS NULL THEN 0 ELSE 1 END AS FollowStatus"
	}

	queryBase := database.DB.
		Select(s).
		Table("follows")

	if userId != "" {
		queryBase = queryBase.Joins(strings.Join([]string{
			"JOIN (",
			"  SELECT",
			"    users.*,",
			"    follows.follow_user_id",
			"  FROM users",
			"    LEFT JOIN (",
			"      SELECT * FROM follows WHERE follower_user_id = ?",
			"    ) AS follows ON follows.follow_user_id = users.id",
			") AS users ON users.id = follower_user_id",
		}, ""), userId)
	} else {
		queryBase = queryBase.Joins("JOIN users ON users.id = follows.follow_user_id")
	}

	queryBase = queryBase.Where("follows.follower_user_id = (SELECT id FROM users WHERE display_id = ? LIMIT 1)", displayId)

	result := queryBase.Limit(limit).Offset(offset).Scan(&items)

	if result.Error != nil {
		return items, count, result.Error
	}

	result = queryBase.Count(&count)

	return items, count, result.Error
}

func (repository FollowRepository) GetFollowersByDisplayId(
	displayId string, userId string, limit int, offset int,
) ([]entity.Follow, int64, error) {
	var items []entity.Follow
	var count int64

	s := strings.Join([]string{
		"users.id AS Id",
		"users.display_id AS DisplayId",
		"users.name AS Name",
		"users.avatar AS Avatar",
	}, ",")

	if userId != "" {
		s = s + ", CASE WHEN users.follow_user_id IS NULL THEN 0 ELSE 1 END AS FollowStatus"
	}

	queryBase := database.DB.
		Select(s).
		Table("follows")

	if userId != "" {
		queryBase = queryBase.Joins(strings.Join([]string{
			"JOIN (",
			"  SELECT",
			"    users.*,",
			"    follows.follow_user_id",
			"  FROM users",
			"    LEFT JOIN (",
			"      SELECT * FROM follows WHERE follower_user_id = ?",
			"    ) AS follows ON follows.follow_user_id = users.id",
			") AS users ON users.id = follower_user_id",
		}, ""), userId)
	} else {
		queryBase = queryBase.Joins("JOIN users ON users.id = follows.follower_user_id")
	}

	queryBase.Where("follows.follow_user_id = (SELECT id FROM users WHERE display_id = ? LIMIT 1)", displayId)

	result := queryBase.Limit(limit).Offset(offset).Scan(&items)

	if result.Error != nil {
		return items, count, result.Error
	}

	result = queryBase.Count(&count)

	return items, count, result.Error
}
