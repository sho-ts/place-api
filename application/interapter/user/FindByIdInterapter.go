package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindByIdInterapter struct {
	UserRepository repository.IUserRepository
}

func NewFindByIdInterapter(
	userRepository repository.IUserRepository,
) FindByIdInterapter {
	return FindByIdInterapter{
		UserRepository: userRepository,
	}
}

func (interapter FindByIdInterapter) Handle(i input.FindByIdInput) (entity.User, error) {
	user, err := interapter.UserRepository.FindById(i.Id)

	return user, err
}
