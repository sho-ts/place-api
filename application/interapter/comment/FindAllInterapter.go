package interapter

import (
	"github.com/sho-ts/place-api/domain/dto/input/comment"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/domain/repository"
)

type FindAllInterapter struct {
	CommentRepository repository.ICommentRepository
}

func NewFindAllInterapter(
	commentRepository repository.ICommentRepository,
) FindAllInterapter {
	return FindAllInterapter{
		CommentRepository: commentRepository,
	}
}

func (interapter FindAllInterapter) Handle(i input.FindAllInput) ([]entity.Comment, error) {
	comments, err := interapter.CommentRepository.FindAll(i.PostId, i.Limit, i.Offset)

	return comments, err
}
