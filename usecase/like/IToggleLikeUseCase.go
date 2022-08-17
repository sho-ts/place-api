package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/like"
)

type IToggleLikeUseCase interface {
	Handle(i input.ToggleLikeInput) error
}
