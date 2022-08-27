package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindByIdInteractor struct {
	UserRepository repository.IUserRepository
}

func NewFindByIdInteractor(
	userRepository repository.IUserRepository,
) FindByIdInteractor {
	return FindByIdInteractor{
		UserRepository: userRepository,
	}
}

func (interactor FindByIdInteractor) Handle(i input.FindByIdInput) (entity.User, error) {
	user, err := interactor.UserRepository.FindById(i.Id)

	return user, err
}
