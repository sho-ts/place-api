package usecase

import (
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/dto/input"
)

type ICreateUserUseCase interface {
	Handle(i input.CreateUserInput) (entity.User, error)
}
