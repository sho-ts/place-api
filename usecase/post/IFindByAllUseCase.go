package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/post"
	"github.com/sho-ts/place-api/domain/entity"
)

type IFindAllUseCase interface {
	Handle(i input.FindAllInput) ([]entity.PostsItem, error)
}
