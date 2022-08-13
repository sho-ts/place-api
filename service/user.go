package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/entity"
)

func CreateUser(authId string, displayId string, name string) (entity.User, error) {
	user := entity.User{
		Id:        authId,
		DisplayId: displayId,
		Name:      name,
	}

	result := database.DB.Create(&user)

	return user, result.Error
}

func GetMe(authId string) (entity.User, error) {
	var user entity.User

	result := database.DB.Where("id = ?", authId).First(&user)

	return user, result.Error
}

func GetUser(userId string) (entity.User, error) {
	var user entity.User

	result := database.DB.Where("display_id = ?", userId).First(&user)

	return user, result.Error
}
