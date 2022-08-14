package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"strings"
)

func CreatePost(postId string, authId string, caption string) (entity.Post, error) {
	post := entity.Post{
		Id:      postId,
		UserId:  authId,
		Caption: caption,
	}

	result := database.DB.Create(&post)

	return post, result.Error
}

func GetPost(postId string) (output.GetPostResponseOutput, error) {
	var post entity.Post
	result := database.DB.Where("id = ?", postId).First(&post)

	var files []entity.Storage

	result = database.DB.Where("post_id = ?", postId).Find(&files)

	o := output.GetPostResponseOutput{
		PostId:  post.Id,
		UserId:  post.UserId,
		Caption: post.Caption,
		Files:   files,
	}

	return o, result.Error
}

func GetPosts(search string, limit int, offset int) ([]output.GetPostsResponseOutput, error) {
	var o []output.GetPostsResponseOutput

	s := strings.Join([]string{
		"posts.id as PostId",
		"posts.user_id as UserId",
		"posts.caption as Caption",
		"storages.url as Thumbnail",
	}, ",")

	// サブクエリで投稿に複数の画像があった場合の重複除外をしている
	j := "join storages on storages.id = (select id from storages s2 where s2.post_id = posts.id limit 1)"

	w := "caption like ?"

	result := database.DB.Table("posts").Select(s).Joins(j).Where(w, "%" + search + "%").Limit(limit).Offset(offset).Scan(&o)

	return o, result.Error
}

func GetUserPosts(userId string, limit int, offset int) ([]output.GetPostsResponseOutput, error) {
	var o []output.GetPostsResponseOutput

	s := strings.Join([]string{
		"posts.id as PostId",
		"posts.user_id as UserId",
		"posts.caption as Caption",
		"storages.url as Thumbnail",
	}, ",")

	w := "posts.user_id = (select id from users where display_id = ?)"

	// サブクエリで投稿に複数の画像があった場合の重複除外をしている
	j := "join storages on storages.id = (select id from storages s2 where s2.post_id = posts.id limit 1)"

	result := database.DB.Table("posts").Select(s).Joins(j).Where(w, userId).Limit(limit).Offset(offset).Scan(&o)

	return o, result.Error
}
