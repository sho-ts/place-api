package interactor

import (
	"github.com/sho-ts/place-api/domain/dto/input/follow"
	"github.com/sho-ts/place-api/domain/dto/output/follow"
	"github.com/sho-ts/place-api/domain/repository"
)

type GetFollowersByDisplayIdInteractor struct {
	followRepository repository.IFollowRepository
}

func NewGetFollowersByDisplayIdInteractor(
	followRepository repository.IFollowRepository,
) GetFollowersByDisplayIdInteractor {
	return GetFollowersByDisplayIdInteractor{
		followRepository: followRepository,
	}
}

func (interactor GetFollowersByDisplayIdInteractor) Handle(i input.GetFollowersByDisplayIdInput) (output.GetFollowersByDisplayIdOutput, error) {
	items, total, err := interactor.followRepository.GetFollowersByDisplayId(i.DisplayId, i.Limit, i.Offset)
	return output.NewGetFollowersByDisplayIdOutput(items, total), err
}
