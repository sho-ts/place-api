package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/infrastructure/database"
	"strings"
)

type UserRepository struct{}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (repository UserRepository) Store(user entity.User) (entity.User, error) {
	result := database.DB.Create(&user)

	return user, result.Error
}

func (repository UserRepository) FindById(userId string) (entity.User, error) {
	user := entity.User{}

	result := database.DB.
		Table("users").
		Select(strings.Join([]string{
			"id AS Id",
			"display_id AS DisplayId",
			"name AS Name",
			"avatar AS Avatar",
		}, ",")).
		Where("id = ?", userId).
		Scan(&user)

	return user, result.Error
}

func (repository UserRepository) FindByDisplayId(displayId string, userId string) (entity.UserDetail, error) {
	user := entity.UserDetail{}

	s := strings.Join([]string{
		"id AS Id",
		"display_id AS DisplayId",
		"name AS Name",
		"avatar AS Avatar",
	}, ",")

	if userId != "" {
		s = s + ",CASE WHEN follows.follow_user_id IS NULL THEN 0 ELSE 1 END AS FollowStatus"
	}

	qb := database.DB.
		Debug().
		Table("users").
		Select(s)

	if userId != "" {
		// ユーザーIDが渡ってきた場合、フォローしているかどうか調べる
		j := "LEFT JOIN (SELECT follow_user_id FROM follows WHERE follow_user_id = (SELECT id FROM users WHERE users.display_id = ? LIMIT 1) AND follower_user_id = ?) AS follows ON follows.follow_user_id = users.id"
		qb = qb.Joins(j, displayId, userId)
	}

	result := qb.
		Where("display_id = ?", displayId).
		Scan(&user)

	return user, result.Error
}

func (repository UserRepository) ChangeProfile(user entity.User) (entity.User, error) {
	result := database.DB.Save(user)

	return user, result.Error
}
