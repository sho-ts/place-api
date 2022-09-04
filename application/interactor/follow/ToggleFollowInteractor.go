package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/follow"
	"github.com/sho-ts/place-api/domain/repository"
)

type ToggleFollowInteractor struct {
	followRepository repository.IFollowRepository
}

func NewToggleFollowInteractor(
	followRepository repository.IFollowRepository,
) ToggleFollowInteractor {
	return ToggleFollowInteractor{
		followRepository: followRepository,
	}
}

func (interactor ToggleFollowInteractor) Handle(i input.ToggleFollowInput) error {
	duplicate, err := interactor.followRepository.CheckDuplicate(i.FollowUserId, i.FollowerUserId)

	if err != nil {
		return err
	}

	if !duplicate {
		err = interactor.followRepository.Store(i.FollowUserId, i.FollowerUserId)
	} else {
		err = interactor.followRepository.Remove(i.FollowUserId, i.FollowerUserId)
	}

	return err
}
