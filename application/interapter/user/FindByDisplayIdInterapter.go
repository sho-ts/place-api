package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindByDisplayIdInterapter struct {
	UserRepository repository.IUserRepository
}

func NewFindByDisplayIdInterapter(
	userRepository repository.IUserRepository,
) FindByDisplayIdInterapter {
	return FindByDisplayIdInterapter{
		UserRepository: userRepository,
	}
}

func (interapter FindByDisplayIdInterapter) Handle(i input.FindByDisplayIdInput) (entity.User, error) {
	user, err := interapter.UserRepository.FindByDisplayId(i.DisplayId)

	return user, err
}
