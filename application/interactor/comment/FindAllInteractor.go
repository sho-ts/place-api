package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/comment"
	output "github.com/sho-ts/place-api/domain/dto/output/comment"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindAllInteractor struct {
	CommentRepository repository.ICommentRepository
}

func NewFindAllInteractor(
	commentRepository repository.ICommentRepository,
) FindAllInteractor {
	return FindAllInteractor{
		CommentRepository: commentRepository,
	}
}

func (interactor FindAllInteractor) Handle(i input.FindAllInput) (output.FindAllOutput, error) {
	comments, err := interactor.CommentRepository.FindAll(i.PostId, i.Limit, i.Offset)
	count, err := interactor.CommentRepository.GetTotalCount(i.PostId)

	o := output.NewFindAllOutput(
		comments,
		count,
	)

	return o, err
}
