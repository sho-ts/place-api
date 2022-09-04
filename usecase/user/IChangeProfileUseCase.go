package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
)

type IChangeProfileUseCase interface {
	Handle(i input.ChangeProfileInput) (entity.User, error)
}
