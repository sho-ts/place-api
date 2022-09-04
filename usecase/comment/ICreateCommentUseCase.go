package usecase

import (
	"github.com/sho-ts/place-api/domain/dto/input/comment"
)

type ICreateCommentUseCase interface {
	Handle(i input.CreateCommentInput) error
}
