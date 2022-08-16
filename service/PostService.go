package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"strings"
	"time"
)

type PostService struct{}

func NewPostService() PostService {
	postService := PostService{}
	return postService
}

func (ps PostService) GetPosts(search string, limit int, offset int) (output.GetPostsOutput, error) {
	var s []struct {
		PostId    string
		Caption   string
		CreatedAt time.Time
		Thumbnail string
		UserId    string
		DisplayId string
		Avatar    string
		Name      string
	}

	// 投稿に複数の画像があった場合の重複除外
	sub := "select id from storages s where s.post_id = posts.id limit 1"

	result := database.DB.
		Table("posts").
		Select(strings.Join([]string{
			"posts.id as PostId",
			"posts.user_id as UserId",
			"posts.caption as Caption",
			"posts.created_at as CreatedAt",
			"storages.url as Thumbnail",
			"users.display_id as DisplayId",
			"users.avatar as Avatar",
			"users.name as Name",
		}, ",")).
		Joins("join storages on storages.id = ("+sub+")").
		Joins("join users on users.id = posts.user_id").
		Where("caption like ?", "%"+search+"%").
		Order("posts.created_at desc").
		Limit(limit).
		Offset(offset).
		Scan(&s)

	items := make([]output.GetPostsOutputItem, len(s))
	for i := 0; i < len(s); i++ {
		items[i] = output.NewGetPostsOutputItem(
			s[i].PostId,
			s[i].Caption,
			s[i].CreatedAt,
			s[i].Thumbnail,
			entity.User{
				DisplayId: s[i].DisplayId,
				Name:      s[i].Name,
				Avatar:    s[i].Avatar,
			},
		)
	}
	o := output.NewGetPostsOutput(items)

	return o, result.Error
}

func (ps PostService) GetUserPosts(userId string, limit int, offset int) (output.GetPostsOutput, error) {
	var s []struct {
		PostId    string
		Caption   string
		CreatedAt time.Time
		Thumbnail string
		UserId    string
		DisplayId string
		Avatar    string
		Name      string
	}

	// 投稿に複数の画像があった場合の重複除外
	sub := "select id from storages s where s.post_id = posts.id limit 1"

	result := database.DB.
		Table("posts").
		Select(strings.Join([]string{
			"posts.id as PostId",
			"posts.user_id as UserId",
			"posts.caption as Caption",
			"posts.created_at as CreatedAt",
			"storages.url as Thumbnail",
			"users.display_id as DisplayId",
			"users.avatar as Avatar",
			"users.name as Name",
		}, ",")).
		Joins("join storages on storages.id = ("+sub+")").
		Joins("join users on users.id = posts.user_id").
		Where("posts.user_id = (select id from users where display_id = ?)", userId).
		Order("posts.created_at desc").
		Limit(limit).
		Offset(offset).
		Scan(&s)

	items := make([]output.GetPostsOutputItem, len(s))
	for i := 0; i < len(s); i++ {
		items[i] = output.NewGetPostsOutputItem(
			s[i].PostId,
			s[i].Caption,
			s[i].CreatedAt,
			s[i].Thumbnail,
			entity.User{
				DisplayId: s[i].DisplayId,
				Name:      s[i].Name,
				Avatar:    s[i].Avatar,
			},
		)
	}
	o := output.NewGetPostsOutput(items)

	return o, result.Error
}
