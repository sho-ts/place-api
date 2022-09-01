package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindByDisplayIdInteractor struct {
	UserRepository repository.IUserRepository
}

func NewFindByDisplayIdInteractor(
	userRepository repository.IUserRepository,
) FindByDisplayIdInteractor {
	return FindByDisplayIdInteractor{
		UserRepository: userRepository,
	}
}

func (interactor FindByDisplayIdInteractor) Handle(i input.FindByDisplayIdInput) (entity.UserDetail, error) {
	user, err := interactor.UserRepository.FindByDisplayId(i.DisplayId, i.UserId)

	return user, err
}
