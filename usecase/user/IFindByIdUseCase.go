package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input"
	"github.com/sho-ts/place-api/domain/entity"
)

type IFindByIdUseCase interface {
	Handle(i input.FindByIdInput) (entity.User, error)
}
