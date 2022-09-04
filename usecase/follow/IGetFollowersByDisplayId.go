package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/follow"
	"github.com/sho-ts/place-api/domain/dto/output/follow"
)

type IGetFollowersByDisplayIdUseCase interface {
	Handle(i input.GetFollowersByDisplayIdInput) (output.GetFollowersByDisplayIdOutput, error)
}
