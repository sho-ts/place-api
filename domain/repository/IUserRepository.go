package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
)

type IUserRepository interface {
	Store(user entity.User) (entity.User, error)
	FindById(userId string) (entity.User, error)
	FindByDisplayId(displayId string, userId string) (entity.UserDetail, error)
	ChangeProfile(user entity.User) (entity.User, error)
}