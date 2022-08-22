package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/follow"
)

type IToggleLikeUseCase interface {
  Handle(i input.ToggleFollowInput) error
}