package interapter

import (
	"github.com/sho-ts/place-api/domain/repository"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/dto/input/user"
)

type UserCreateInterapter struct {
	UserRepository repository.IUserRepository
}

func NewUserCreateInterapter(
	userRepository repository.IUserRepository,
) UserCreateInterapter {
	return UserCreateInterapter{
		UserRepository: userRepository,
	}
}

func (interapter UserCreateInterapter) Handle(i input.CreateUserInput) (entity.User, error) {
	user := entity.NewUser(
		i.UserId,
		i.DisplayId,
		i.Name,
	)

	user, err := interapter.UserRepository.Store(user)

	return user, err
}