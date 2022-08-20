package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type ChangeProfileInterapter struct {
	UserRepository    repository.IUserRepository
	StorageRepository repository.IStorageRepository
}

func NewChangeProfileInterapter(
	userRepository repository.IUserRepository,
	storageRepository repository.IStorageRepository,
) ChangeProfileInterapter {
	return ChangeProfileInterapter{
		UserRepository:    userRepository,
		StorageRepository: storageRepository,
	}
}

func (interapter ChangeProfileInterapter) Handle(i input.ChangeProfileInput) (entity.User, error) {
	var path string
	var err error

	if i.FileName != "" && i.File != nil {
		path, err = interapter.StorageRepository.UploadToS3Bucket(i.File, i.FileName)
	}

	prevUser, err := interapter.UserRepository.FindById(i.UserId)

	draftUser := entity.User{
		Id:        i.UserId,
		DisplayId: i.DisplayId,
		Name:      i.Name,
		Avatar:    path,
	}

  if i.DisplayId == "" {
    draftUser.DisplayId = prevUser.DisplayId
  }

  if i.Name == "" {
    draftUser.Name = prevUser.Name
  }

	if path == "" {
		draftUser.Avatar = prevUser.Avatar
	}

	user, err := interapter.UserRepository.ChangeProfile(draftUser)

	return user, err
}
