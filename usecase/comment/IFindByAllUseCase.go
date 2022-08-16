package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/comment"
	"github.com/sho-ts/place-api/domain/entity"
)

type IFindAllUseCase interface {
	Handle(i input.FindAllInput) ([]entity.Comment, error)
}
