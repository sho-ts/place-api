package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/like"
	"github.com/sho-ts/place-api/domain/repository"
)

type ToggleLikeInteractor struct {
	LikeRepository repository.ILikeRepository
}

func NewToggleLikeInteractor(
	likeRepository repository.ILikeRepository,
) ToggleLikeInteractor {
	return ToggleLikeInteractor{
		LikeRepository: likeRepository,
	}
}

func (interactor ToggleLikeInteractor) Handle(i input.ToggleLikeInput) error {
	duplicate, err := interactor.LikeRepository.CheckDuplicate(i.PostId, i.UserId)

	if err != nil {
		return err
	}

	if duplicate {
		err = interactor.LikeRepository.Remove(i.PostId, i.UserId)
	} else {
		err = interactor.LikeRepository.Store(i.PostId, i.UserId)
	}

	return err
}
