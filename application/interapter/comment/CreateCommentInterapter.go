package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/comment"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type CreateCommentInterapter struct {
	CommentRepository repository.ICommentRepository
}

func NewCreateCommentInterapter(
	commentRepository repository.ICommentRepository,
) CreateCommentInterapter {
	return CreateCommentInterapter{
		CommentRepository: commentRepository,
	}
}

func (interapter CreateCommentInterapter) Handle(i input.CreateCommentInput) error {
	comment := entity.NewComment(
		i.CommentId,
		i.PostId,
		i.Content,
		entity.User{Id: i.UserId},
	)

	return interapter.CommentRepository.Store(comment)
}
