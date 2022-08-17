package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/comment"
	"github.com/sho-ts/place-api/domain/dto/output/comment"
)

type IFindAllUseCase interface {
	Handle(i input.FindAllInput) (output.FindAllOutput, error)
}
