package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/follow"
	"github.com/sho-ts/place-api/domain/dto/output/follow"
)

type IGetFollowsByDisplayIdUseCase interface {
	Handle(i input.GetFollowsByDisplayIdInput) (output.GetFollowsByDisplayIdOutput, error)
}
