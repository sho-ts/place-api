package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type ChangeProfileInteractor struct {
	UserRepository    repository.IUserRepository
	StorageRepository repository.IStorageRepository
}

func NewChangeProfileInteractor(
	userRepository repository.IUserRepository,
	storageRepository repository.IStorageRepository,
) ChangeProfileInteractor {
	return ChangeProfileInteractor{
		UserRepository:    userRepository,
		StorageRepository: storageRepository,
	}
}

func (interactor ChangeProfileInteractor) Handle(i input.ChangeProfileInput) (entity.User, error) {
	var path string
	var err error

	if i.FileName != "" && i.File != nil {
		path, err = interactor.StorageRepository.UploadToS3Bucket(i.File, i.FileName)
	}

	prevUser, err := interactor.UserRepository.FindById(i.UserId)

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

	user, err := interactor.UserRepository.ChangeProfile(draftUser)

	return user, err
}
