package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/post"
)

type ICreatePostUseCase interface {
	Handle(i input.CreatePostInput) error
}
