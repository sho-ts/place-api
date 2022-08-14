package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"strings"
	"time"
)

func CreateComment(i input.CreateCommentInput) (entity.Comment, error) {
	comment := entity.Comment{
		Id:      i.Id,
		UserId:  i.UserId,
		PostId:  i.PostId,
		Content: i.Content,
	}

	result := database.DB.Create(&comment)

	return comment, result.Error
}

func GetComments(postId string, limit int, offset int) ([]output.GetCommentOutput, error) {
	var s []struct {
		CommentId string
		Content   string
		PostId    string
		CreatedAt time.Time
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
			"users.display_id as DisplayId",
			"users.avatar as Avatar",
			"users.name as UserName",
		}, ",")).
		Joins("join users on users.id = comments.user_id").
		Where("comments.post_id = ?", postId).
		Order("comments.created_at desc").
		Limit(limit).
		Offset(offset).
		Scan(&s)

	o := make([]output.GetCommentOutput, len(s))
	for i := 0; i < len(s); i++ {
		o[i] = output.GetCommentOutput{
			CommentId: s[i].CommentId,
			PostId:    s[i].PostId,
			Content:   s[i].Content,
			CreatedAt: s[i].CreatedAt,
			User: entity.User{
				DisplayId: s[i].DisplayId,
				Name:      s[i].UserName,
				Avatar:    s[i].Avatar,
			},
		}
	}

	return o, result.Error
}
