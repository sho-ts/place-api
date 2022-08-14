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
	var post entity.Post
	result := database.DB.Where("id = ?", postId).First(&post)

	var files []entity.Storage

	result = database.DB.Where("post_id = ?", postId).Find(&files)

	o := output.GetPostOutput{
		PostId:  post.Id,
		UserId:  post.UserId,
		Caption: post.Caption,
		Files:   files,
	}

	return o, result.Error
}

func GetPosts(search string, limit int, offset int) ([]output.GetPostsOutput, error) {
	var o []output.GetPostsOutput

	s := strings.Join([]string{
		"posts.id as PostId",
		"posts.user_id as UserId",
		"posts.caption as Caption",
		"storages.url as Thumbnail",
	}, ",")

	w := "caption like ?"

	// 投稿に複数の画像があった場合の重複除外
	sq := "select id from storages s2 where s2.post_id = posts.id limit 1"

	j := "join storages on storages.id = (" + sq + ")"

	result := database.DB.Table("posts").Select(s).Joins(j).Where(w, "%"+search+"%").Limit(limit).Offset(offset).Scan(&o)

	return o, result.Error
}

func GetUserPosts(userId string, limit int, offset int) ([]output.GetPostsOutput, error) {
	var o []output.GetPostsOutput

	s := strings.Join([]string{
		"posts.id as PostId",
		"posts.user_id as UserId",
		"posts.caption as Caption",
		"storages.url as Thumbnail",
	}, ",")

	w := "posts.user_id = (select id from users where display_id = ?)"

	// 投稿に複数の画像があった場合の重複除外
	sq := "select id from storages s2 where s2.post_id = posts.id limit 1"

	j := "join storages on storages.id = (" + sq + ")"

	result := database.DB.Table("posts").Select(s).Joins(j).Where(w, userId).Limit(limit).Offset(offset).Scan(&o)

	return o, result.Error
}
