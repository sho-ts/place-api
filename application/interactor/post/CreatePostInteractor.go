package interactor

import (
	"github.com/sho-ts/place-api/application/util"
	input "github.com/sho-ts/place-api/domain/dto/input/post"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type CreatePostInteractor struct {
	PostRepository    repository.IPostRepository
	StorageRepository repository.IStorageRepository
}

func NewCreatePostInteractor(
	postRepository repository.IPostRepository,
	storageRepository repository.IStorageRepository,
) CreatePostInteractor {
	return CreatePostInteractor{
		PostRepository:    postRepository,
		StorageRepository: storageRepository,
	}
}

func (interactor CreatePostInteractor) Handle(i input.CreatePostInput) error {
	path, err := interactor.StorageRepository.UploadToS3Bucket(i.File, i.FileName)

	storageObjects := []entity.StorageObject{
		entity.NewStorageObject(
			util.GetULID(),
			i.PostId,
			i.UserId,
			path,
		),
	}

	post := entity.NewPost(
		i.PostId,
		i.Caption,
		0,
		storageObjects,
		entity.NewUser(i.UserId, "", ""),
	)

	_, err = interactor.PostRepository.Store(post)

	return err
}
