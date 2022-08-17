package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/post"
	"github.com/sho-ts/place-api/domain/dto/output/post"
)

type IFindAllUseCase interface {
	Handle(i input.FindAllInput) (output.FindAllOutput, error)
}
