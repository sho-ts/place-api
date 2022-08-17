package repository

import (
	"strings"
	"time"

	"github.com/sho-ts/place-api/application/util"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/infrastructure/database/table"
	"gorm.io/gorm"
)

type PostRepository struct{}

func NewPostRepository() PostRepository {
	return PostRepository{}
}

func (repository PostRepository) Store(post entity.Post) (entity.Post, error) {
	tx := database.DB.Begin()

	postData := table.Post{
		Id:      post.PostId,
		UserId:  post.User.Id,
		Caption: post.Caption,
	}

	result := database.DB.Create(&postData)

	storage := table.Storage{
		Id:     util.GetULID(),
		UserId: post.User.Id,
		PostId: post.PostId,
		Url:    post.StorageObjects[0].Url,
	}

	result = database.DB.Create(&storage)

	if result.Error != nil {
		tx.Rollback()
		return post, result.Error
	}

	tx.Commit()

	return post, result.Error
}

func (repository PostRepository) FindById(postId string, userId string) (entity.Post, error) {
	var postResult struct {
		PostId    string
		Caption   string
		CreatedAt time.Time
		UserId    string
		DisplayId string
		Avatar    string
		Name      string
		Liked     int
	}

	s := strings.Join([]string{
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
		s = s + ",case when liked.post_id is null then 0 else 1 end as Liked"
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
		Table("posts").
		Select(s).
		Joins(sj, sub).
		Joins("join users on users.id = posts.user_id").
		Where("posts.id = ?", postId).
		Scan(&postResult)

	var storageObjects []entity.StorageObject

	result = database.DB.
		Table("storages").
		Select([]string{
			"id as Id",
			"post_id as PostId",
			"user_id as UserId",
			"url as Url",
		}).
		Where("post_id = ?", postId).
		Find(&storageObjects)

	return entity.Post{
		PostId:         postResult.PostId,
		Caption:        postResult.Caption,
		CreatedAt:      postResult.CreatedAt,
		Liked:          postResult.Liked,
		StorageObjects: storageObjects,
		User: entity.User{
			Id:        postResult.UserId,
			DisplayId: postResult.DisplayId,
			Name:      postResult.Name,
			Avatar:    postResult.Avatar,
		},
	}, result.Error
}

func (repository PostRepository) FindAll(userId string, limit int, offset int) ([]entity.PostsItem, error) {
	var postsResult []struct {
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

	var result *gorm.DB

	base := database.DB.
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
		Joins("join storages on storages.id = (" + sub + ")").
		Joins("join users on users.id = posts.user_id")

	if userId != "" {
		result = base.Where("posts.user_id = (select id from users where display_id = ?)", userId).
			Order("posts.created_at desc").
			Limit(limit).
			Offset(offset).
			Scan(&postsResult)
	} else {
		result = base.
			Order("posts.created_at desc").
			Limit(limit).
			Offset(offset).
			Scan(&postsResult)
	}

	items := make([]entity.PostsItem, len(postsResult))
	for index := range postsResult {
		items[index] = entity.PostsItem{
			PostId:    postsResult[index].PostId,
			Caption:   postsResult[index].Caption,
			CreatedAt: postsResult[index].CreatedAt,
			Thumbnail: postsResult[index].Thumbnail,
			User: entity.User{
				Id:        postsResult[index].UserId,
				DisplayId: postsResult[index].DisplayId,
				Name:      postsResult[index].Name,
				Avatar:    postsResult[index].Avatar,
			},
		}
	}

	return items, result.Error
}

func (repository PostRepository) GetTotalCount(userId string) (int64, error) {
	var count int64
	var result *gorm.DB

	if userId != "" {
		result = database.DB.
			Table("posts").
			Where("posts.user_id = (select id from users where display_id = ?)", userId).
			Count(&count)
	} else {
		result = database.DB.
			Table("posts").
			Count(&count)
	}

	return count, result.Error
}
