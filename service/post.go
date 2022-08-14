package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
	"strings"
)

func CreatePost(i input.CreatePostInput) (entity.Post, error) {
	tx := database.DB.Begin()

	post := entity.Post{
		Id:      i.PostId,
		UserId:  i.UserId,
		Caption: i.Caption,
	}

	result := database.DB.Create(&post)

	storage := entity.Storage{
		Id:     util.GetULID(),
		UserId: i.UserId,
		PostId: i.PostId,
		Url:    i.Urls[0],
	}

	result = database.DB.Create(&storage)

	if result.Error != nil {
		tx.Rollback()
		return post, result.Error
	}

	tx.Commit()

	return post, result.Error
}

func GetPost(postId string) (output.GetPostOutput, error) {
	var post struct {
		PostId    string `json:"postId"`
		Caption   string `json:"caption"`
		UserId    string `json:"userId"`
		DisplayId string `json:"displayId"`
		Avatar    string `json:"avatar"`
		Name      string `json:"name"`
	}

	s := strings.Join([]string{
		"posts.id as PostId",
		"posts.caption as Caption",
		"posts.user_id as UserId",
		"users.display_id as DisplayId",
		"users.avatar as Avatar",
		"users.name as Name",
	}, ",")

	result := database.DB.Table("posts").
		Select(s).
		Joins("join users on users.id = posts.user_id").
		Where("posts.id = ?", postId).
		Scan(&post)

	var files []entity.Storage

	result = database.DB.Where("post_id = ?", postId).Find(&files)

	o := output.GetPostOutput{
		PostId:  post.PostId,
		Caption: post.Caption,
		Files:   files,
		User: entity.User{
			Id:        post.UserId,
			DisplayId: post.DisplayId,
			Avatar:    post.Avatar,
			Name:      post.Name,
		},
	}

	return o, result.Error
}

func GetPosts(search string, limit int, offset int) ([]output.GetPostsOutput, error) {
	var posts []struct {
		PostId    string `json:"postId"`
		Caption   string `json:"caption"`
		Thumbnail string `json:"thumbnail"`
		UserId    string `json:"userId"`
		DisplayId string `json:"displayId"`
		Avatar    string `json:"avatar"`
		Name      string `json:"name"`
	}

	s := strings.Join([]string{
		"posts.id as PostId",
		"posts.user_id as UserId",
		"posts.caption as Caption",
		"storages.url as Thumbnail",
		"users.display_id as DisplayId",
		"users.avatar as Avatar",
		"users.name as Name",
	}, ",")

	// 投稿に複数の画像があった場合の重複除外
	sq := "select id from storages s2 where s2.post_id = posts.id limit 1"

	result := database.DB.
		Table("posts").
		Select(s).
		Joins("join storages on storages.id = ("+sq+")").
		Joins("join users on users.id = posts.user_id").
		Where("caption like ?", "%"+search+"%").
		Limit(limit).
		Offset(offset).
		Scan(&posts)

	o := make([]output.GetPostsOutput, len(posts))
	for i := 0; i < len(posts); i++ {
		o[i] = output.GetPostsOutput{
			PostId:    posts[i].PostId,
			Caption:   posts[i].Caption,
			Thumbnail: posts[i].Thumbnail,
			User: entity.User{
				DisplayId: posts[i].DisplayId,
				Name:      posts[i].Name,
				Avatar:    posts[i].Avatar,
			},
		}
	}

	return o, result.Error
}

func GetUserPosts(userId string, limit int, offset int) ([]output.GetPostsOutput, error) {
	var posts []struct {
		PostId    string `json:"postId"`
		Caption   string `json:"caption"`
		Thumbnail string `json:"thumbnail"`
		UserId    string `json:"userId"`
		DisplayId string `json:"displayId"`
		Avatar    string `json:"avatar"`
		Name      string `json:"name"`
	}

	s := strings.Join([]string{
		"posts.id as PostId",
		"posts.user_id as UserId",
		"posts.caption as Caption",
		"storages.url as Thumbnail",
		"users.display_id as DisplayId",
		"users.avatar as Avatar",
		"users.name as Name",
	}, ",")

	// 投稿に複数の画像があった場合の重複除外
	sq := "select id from storages s2 where s2.post_id = posts.id limit 1"

	result := database.DB.
		Table("posts").
		Select(s).
		Joins("join storages on storages.id = ("+sq+")").
		Joins("join users on users.id = posts.user_id").
		Where("posts.user_id = (select id from users where display_id = ?)", userId).
		Limit(limit).
		Offset(offset).
		Scan(&posts)

	o := make([]output.GetPostsOutput, len(posts))
	for i := 0; i < len(posts); i++ {
		o[i] = output.GetPostsOutput{
			PostId:    posts[i].PostId,
			Caption:   posts[i].Caption,
			Thumbnail: posts[i].Thumbnail,
			User: entity.User{
				DisplayId: posts[i].DisplayId,
				Name:      posts[i].Name,
				Avatar:    posts[i].Avatar,
			},
		}
	}

	return o, result.Error
}
