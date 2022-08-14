package service

import (
	"strings"

	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
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

func GetComments(postId string, limit int, offset int) ([]output.GetCommentsOutput, error) {
	var o []output.GetCommentsOutput

	s := strings.Join([]string{
		"comments.id as Id",
		"comments.content as Content",
		"comments.post_id as PostId",
		"users.display_id as UserId",
		"users.avatar as Avatar",
		"users.name as UserName",
	}, ",")

	w := "comments.post_id = ?"

	j := "join users on users.id = comments.user_id"

	result := database.DB.Table("comments").Select(s).Joins(j).Where(w, postId).Limit(limit).Offset(offset).Scan(&o)

	return o, result.Error
}
