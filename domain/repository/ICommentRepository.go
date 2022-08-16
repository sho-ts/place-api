package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
)

type ICommentRepository interface {
	Store(comment entity.Comment) error
}
