package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/post"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
	"github.com/sho-ts/place-api/util"
)

type CreatePostInterapter struct {
	PostRepository    repository.IPostRepository
	StorageRepository repository.IStorageRepository
}

func NewCreatePostInterapter(
	postRepository repository.IPostRepository,
	storageRepository repository.IStorageRepository,
) CreatePostInterapter {
	return CreatePostInterapter{
		PostRepository:    postRepository,
		StorageRepository: storageRepository,
	}
}

func (interapter CreatePostInterapter) Handle(i input.CreatePostInput) (error) {
	path, err := interapter.StorageRepository.UploadToS3Bucket(i.File, i.FileName)

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

	_, err = interapter.PostRepository.Store(post)

	return err
}
