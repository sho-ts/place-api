package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
)

type IPostRepository interface {
	Store(entity.Post) (entity.Post, error)
	// FindById() (entity.Post, error)
	// FindAll() ([]entity.PostsItem, error)
	// Count() (int, error)
}
