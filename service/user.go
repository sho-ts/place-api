package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/entity"
)

type UserService struct{}

func NewUserService() UserService {
	userService := UserService{}
	return userService
}

func (this UserService) CreateUser(i input.CreateUserInput) (entity.User, error) {
	user := entity.User{
		Id:        i.UserId,
		DisplayId: i.DisplayId,
		Name:      i.Name,
	}

	result := database.DB.Create(&user)

	return user, result.Error
}

func (this UserService) GetMe(authId string) (entity.User, error) {
	var user entity.User

	result := database.DB.Where("id = ?", authId).First(&user)

	return user, result.Error
}

func (this UserService) GetUser(userId string) (entity.User, error) {
	var user entity.User

	result := database.DB.Where("display_id = ?", userId).First(&user)

	return user, result.Error
}