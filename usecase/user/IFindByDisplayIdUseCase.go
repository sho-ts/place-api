package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input"
	"github.com/sho-ts/place-api/domain/entity"
)

type IFindByDisplayIdUseCase interface {
	Handle(i input.FindByDisplayIdInput) (entity.User, error)
}
