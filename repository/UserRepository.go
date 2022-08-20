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
			"id as Id",
			"display_id as DisplayId",
			"name as Name",
			"avatar as Avatar",
		}, ",")).
		Where("id = ?", userId).
		Scan(&user)

	return user, result.Error
}

func (repository UserRepository) FindByDisplayId(displayId string) (entity.User, error) {
	user := entity.User{}

	result := database.DB.
		Table("users").
		Select(strings.Join([]string{
			"id as Id",
			"display_id as DisplayId",
			"name as Name",
			"avatar as Avatar",
		}, ",")).
		Where("display_id = ?", displayId).
		Scan(&user)

	return user, result.Error
}
