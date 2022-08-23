package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/follow"
	"github.com/sho-ts/place-api/domain/repository"
)

type ToggleFollowInterapter struct {
	followRepository repository.IFollowRepository
}

func NewToggleFollowInterapter(
	followRepository repository.IFollowRepository,
) ToggleFollowInterapter {
	return ToggleFollowInterapter{
		followRepository: followRepository,
	}
}

func (interapter ToggleFollowInterapter) Handle(i input.ToggleFollowInput) error {
	duplicate, err := interapter.followRepository.CheckDuplicate(i.FollowUserId, i.FollowerUserId)

  if err != nil {
    return err
  }

  if !duplicate {
    err = interapter.followRepository.Store(i.FollowUserId,i.FollowerUserId)
  } else {
    err = interapter.followRepository.Remove(i.FollowUserId,i.FollowerUserId)
  }

  return err
}
