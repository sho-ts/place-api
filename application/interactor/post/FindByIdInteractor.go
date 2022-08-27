package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/post"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindByIdInteractor struct {
	PostRepository repository.IPostRepository
}

func NewFindByIdInteractor(
	postRepository repository.IPostRepository,
) FindByIdInteractor {
	return FindByIdInteractor{
		PostRepository: postRepository,
	}
}

func (interactor FindByIdInteractor) Handle(i input.FindByIdInput) (entity.Post, error) {
	post, err := interactor.PostRepository.FindById(i.PostId, i.UserId)

	return post, err
}
