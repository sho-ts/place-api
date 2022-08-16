package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/infrastructure/database/table"
)

type CommentRepository struct{}

func NewCommentRepository() CommentRepository {
	return CommentRepository{}
}

func (repository CommentRepository) Store(comment entity.Comment) error {
	commentData := table.Comment{
		Id:      comment.CommentId,
		UserId:  comment.User.Id,
		PostId:  comment.PostId,
		Content: comment.Content,
	}

	result := database.DB.Create(&commentData)

	return result.Error
}
