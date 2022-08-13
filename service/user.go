package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
)

func CreateUser(authId string, displayId string, name string) entity.User {
	user := entity.User{
		Id:        util.GetULID(),
		AuthId:    authId,
		DisplayId: displayId,
		Name:      name,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		panic("何かがおかしい")
	}

	return user
}

func GetMe(authId string) entity.User {
	var user entity.User

	database.DB.Where("auth_id = ?", authId).First(&user)

	return user
}
