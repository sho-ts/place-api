package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type UserCreateInteractor struct {
	UserRepository repository.IUserRepository
}

func NewUserCreateInteractor(
	userRepository repository.IUserRepository,
) UserCreateInteractor {
	return UserCreateInteractor{
		UserRepository: userRepository,
	}
}

func (interactor UserCreateInteractor) Handle(i input.CreateUserInput) (entity.User, error) {
	user := entity.NewUser(
		i.UserId,
		i.DisplayId,
		i.Name,
	)

	user, err := interactor.UserRepository.Store(user)

	return user, err
}
