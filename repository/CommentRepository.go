package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/infrastructure/database/table"
	"strings"
	"time"
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

func (repository CommentRepository) FindAll(postId string, limit int, offset int) (
	[]entity.Comment, error,
) {
	var commentsResult []struct {
		CommentId string
		Content   string
		PostId    string
		CreatedAt time.Time
		UserId    string
		DisplayId string
		Avatar    string
		UserName  string
	}

	result := database.DB.
		Table("comments").
		Select(strings.Join([]string{
			"comments.id as CommentId",
			"comments.content as Content",
			"comments.post_id as PostId",
			"comments.created_at as CreatedAt",
			"users.id as UserId",
			"users.display_id as DisplayId",
			"users.avatar as Avatar",
			"users.name as UserName",
		}, ",")).
		Joins("join users on users.id = comments.user_id").
		Where("comments.post_id = ?", postId).
		Order("comments.created_at desc").
		Limit(limit).
		Offset(offset).
		Scan(&commentsResult)

	items := make([]entity.Comment, len(commentsResult))
	for index := range commentsResult {
		items[index] = entity.Comment{
			CommentId: commentsResult[index].CommentId,
			PostId:    commentsResult[index].PostId,
			Content:   commentsResult[index].Content,
			CreatedAt: commentsResult[index].CreatedAt,
			User: entity.User{
				Id:        commentsResult[index].UserId,
				DisplayId: commentsResult[index].DisplayId,
				Name:      commentsResult[index].UserName,
				Avatar:    commentsResult[index].Avatar,
			},
		}
	}

	return items, result.Error
}

func (repository CommentRepository) GetTotalCount(postId string) (int64, error) {
	var count int64

	result := database.DB.
		Table("comments").
		Where("comments.post_id = ?", postId).
		Count(&count)

	return count, result.Error
}
