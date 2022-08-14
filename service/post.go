package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
	"strings"
	"time"
)

/* 投稿を新規作成する */
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

/* 指定したIDの投稿を取得する */
func GetPost(postId string, userId string) (output.GetPostOutput, error) {
	var s struct {
		PostId    string
		Caption   string
		CreatedAt time.Time
		UserId    string
		DisplayId string
		Avatar    string
		Name      string
		Liked     int
	}

	sel := strings.Join([]string{
		"posts.id as PostId",
		"posts.caption as Caption",
		"posts.user_id as UserId",
		"posts.created_at as CreatedAt",
		"users.display_id as DisplayId",
		"users.avatar as Avatar",
		"users.name as Name",
	}, ",")

	// ------------------------
	// ユーザーIDを渡している場合、既にいいねしているかどうかも調べる
	if userId != "" {
		sel = sel + ",case when liked.post_id is null then 0 else 1 end as Liked"
	}

	// サブクエリ
	sub := database.DB.Table("likes").
		Select("post_id").
		Where("user_id = ?", userId).
		Where("post_id = ?", postId).
		Limit(1)

	var sj string
	if userId != "" {
		sj = "left join (?) as liked on liked.post_id = posts.id"
	}
	// ------------------------

	result := database.DB.
		Debug().
		Table("posts").
		Select(sel).
		Joins(sj, sub).
		Joins("join users on users.id = posts.user_id").
		Where("posts.id = ?", postId).
		Scan(&s)

	var files []entity.Storage

	result = database.DB.
		Where("post_id = ?", postId).
		Find(&files)

	o := output.GetPostOutput{
		PostId:    s.PostId,
		Caption:   s.Caption,
		CreatedAt: s.CreatedAt,
		Liked:     s.Liked,
		Files:     files,
		User: entity.User{
			Id:        s.UserId,
			DisplayId: s.DisplayId,
			Avatar:    s.Avatar,
			Name:      s.Name,
		},
	}

	return o, result.Error
}

/* 全てのユーザーの投稿一覧を取得する */
func GetPosts(search string, limit int, offset int) ([]output.GetPostsOutput, error) {
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

	o := make([]output.GetPostsOutput, len(s))
	for i := 0; i < len(s); i++ {
		o[i] = output.GetPostsOutput{
			PostId:    s[i].PostId,
			Caption:   s[i].Caption,
			CreatedAt: s[i].CreatedAt,
			Thumbnail: s[i].Thumbnail,
			User: entity.User{
				DisplayId: s[i].DisplayId,
				Name:      s[i].Name,
				Avatar:    s[i].Avatar,
			},
		}
	}

	return o, result.Error
}

/* 指定したユーザーの投稿一覧を取得する */
func GetUserPosts(userId string, limit int, offset int) ([]output.GetPostsOutput, error) {
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

	o := make([]output.GetPostsOutput, len(s))
	for i := 0; i < len(s); i++ {
		o[i] = output.GetPostsOutput{
			PostId:    s[i].PostId,
			Caption:   s[i].Caption,
			CreatedAt: s[i].CreatedAt,
			Thumbnail: s[i].Thumbnail,
			User: entity.User{
				DisplayId: s[i].DisplayId,
				Name:      s[i].Name,
				Avatar:    s[i].Avatar,
			},
		}
	}

	return o, result.Error
}
