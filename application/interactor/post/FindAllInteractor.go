package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/post"
	output "github.com/sho-ts/place-api/domain/dto/output/post"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindAllInteractor struct {
	PostRepository repository.IPostRepository
}

func NewFindAllInteractor(
	postRepository repository.IPostRepository,
) FindAllInteractor {
	return FindAllInteractor{
		PostRepository: postRepository,
	}
}

func (interactor FindAllInteractor) Handle(i input.FindAllInput) (output.FindAllOutput, error) {
	posts, err := interactor.PostRepository.FindAll(i.UserId, i.Limit, i.Offset)
	count, err := interactor.PostRepository.GetTotalCount(i.UserId)

	o := output.NewFindAllOutput(
		posts,
		count,
	)

	return o, err
}
