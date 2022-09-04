package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
)

type ICreateUserUseCase interface {
	Handle(i input.CreateUserInput) (entity.User, error)
}
