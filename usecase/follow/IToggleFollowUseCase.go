package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/follow"
)

type IToggleFollowUseCase interface {
  Handle(i input.ToggleFollowInput) error
}