package usecase

import (
  "github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/dto/input/user"
)

type IFindByDisplayIdUseCase interface {
	Handle(i input.FindByDisplayIdInput) (entity.UserDetail, error)
}
