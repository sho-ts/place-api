package interactor

import (
	input "github.com/sho-ts/place-api/domain/dto/input/comment"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type CreateCommentInteractor struct {
	CommentRepository repository.ICommentRepository
}

func NewCreateCommentInteractor(
	commentRepository repository.ICommentRepository,
) CreateCommentInteractor {
	return CreateCommentInteractor{
		CommentRepository: commentRepository,
	}
}

func (interactor CreateCommentInteractor) Handle(i input.CreateCommentInput) error {
	comment := entity.NewComment(
		i.CommentId,
		i.PostId,
		i.Content,
		entity.User{Id: i.UserId},
	)

	return interactor.CommentRepository.Store(comment)
}
