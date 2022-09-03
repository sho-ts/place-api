package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
)

type IPostRepository interface {
	Store(entity.Post) (entity.Post, error)
	FindById(postId string, userId string) (entity.Post, error)
	FindAll(displayId string, search string, limit int, offset int) ([]entity.PostsItem, error)
	GetTotalCount(displayId string, search string) (int64, error)
}
