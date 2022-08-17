package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/post"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindByIdInterapter struct {
	PostRepository repository.IPostRepository
}

func NewFindByIdInterapter(
	postRepository repository.IPostRepository,
) FindByIdInterapter {
	return FindByIdInterapter{
		PostRepository: postRepository,
	}
}

func (interapter FindByIdInterapter) Handle(i input.FindByIdInput) (entity.Post, error) {
	post, err := interapter.PostRepository.FindById(i.PostId, i.UserId)

	return post, err
}
