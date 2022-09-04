package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
)

type IFindByDisplayIdUseCase interface {
	Handle(i input.FindByDisplayIdInput) (entity.UserDetail, error)
}
