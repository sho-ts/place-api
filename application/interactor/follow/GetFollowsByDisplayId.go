package interactor

import (
	"github.com/sho-ts/place-api/domain/dto/input/follow"
	"github.com/sho-ts/place-api/domain/dto/output/follow"
	"github.com/sho-ts/place-api/domain/repository"
)

type GetFollowsByDisplayIdInteractor struct {
	followRepository repository.IFollowRepository
}

func NewGetFollowsByDisplayIdInteractor(
	followRepository repository.IFollowRepository,
) GetFollowsByDisplayIdInteractor {
	return GetFollowsByDisplayIdInteractor{
		followRepository: followRepository,
	}
}

func (interactor GetFollowsByDisplayIdInteractor) Handle(i input.GetFollowsByDisplayIdInput) (output.GetFollowsByDisplayIdOutput, error) {
	items, total, err := interactor.followRepository.GetFollowsByDisplayId(i.DisplayId, i.UserId, i.Limit, i.Offset)
	return output.NewGetFollowsByDisplayIdOutput(items, total), err
}
