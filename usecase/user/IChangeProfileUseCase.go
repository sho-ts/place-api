package usecase

import (
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/dto/input/user"
)

type IChangeProfileUseCase interface {
	Handle(i input.ChangeProfileInput) (entity.User, error)
}
