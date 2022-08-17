package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/like"
	"github.com/sho-ts/place-api/domain/repository"
)

type ToggleLikeInterapter struct {
	LikeRepository repository.ILikeRepository
}

func NewToggleLikeInterapter(
	likeRepository repository.ILikeRepository,
) ToggleLikeInterapter {
	return ToggleLikeInterapter{
		LikeRepository: likeRepository,
	}
}

func (interapter ToggleLikeInterapter) Handle(i input.ToggleLikeInput) error {
	duplicate, err := interapter.LikeRepository.CheckDuplicate(i.PostId, i.UserId)

	if duplicate {
		err = interapter.LikeRepository.Remove(i.PostId, i.UserId)
	} else {
		err = interapter.LikeRepository.Store(i.PostId, i.UserId)
	}

	return err
}
