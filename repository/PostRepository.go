package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/infrastructure/database/table"
	"github.com/sho-ts/place-api/util"
)

type PostRepository struct{}

func NewPostRepository() PostRepository {
	return PostRepository{}
}

func (repository PostRepository) Store(post entity.Post) (entity.Post, error) {
	tx := database.DB.Begin()

	postData := table.Post{
		Id:      post.PostId,
		UserId:  post.User.Id,
		Caption: post.Caption,
	}

	result := database.DB.Create(&postData)

	storage := table.Storage{
		Id:     util.GetULID(),
		UserId: post.User.Id,
		PostId: post.PostId,
		Url:    post.StorageObjects[0].Url,
	}

	result = database.DB.Create(&storage)

	if result.Error != nil {
		tx.Rollback()
		return post, result.Error
	}

	tx.Commit()

	return post, result.Error
}
