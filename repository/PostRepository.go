package repository

import (
	"github.com/sho-ts/place-api/application/util"
	"github.com/sho-ts/place-api/domain/entity"
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/infrastructure/database/table"
	"strings"
	"time"
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

	if result.Error != nil {
		return post, result.Error
	}

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
		"posts.id AS PostId",
		"posts.caption AS Caption",
		"posts.user_id AS UserId",
		"posts.created_at AS CreatedAt",
		"users.display_id AS DisplayId",
		"users.avatar AS Avatar",
		"users.name AS Name",
	}, ",")

	// ------------------------
	// ユーザーIDを渡している場合、既にいいねしているかどうかも調べる
	if userId != "" {
		s = s + ",CASE WHEN liked.post_id IS NULL THEN 0 ELSE 1 END AS Liked"
	}

	// サブクエリ
	sub := database.DB.Table("likes").
		Select("post_id").
		Where("user_id = ?", userId).
		Where("post_id = ?", postId).
		Limit(1)

	var sj string
	if userId != "" {
		sj = "LEFT JOIN (?) AS liked ON liked.post_id = posts.id"
	}
	// ------------------------

	result := database.DB.
		Table("posts").
		Select(s).
		Joins(sj, sub).
		Joins("JOIN users ON users.id = posts.user_id").
		Where("posts.id = ?", postId).
		Scan(&postResult)

	if result.Error != nil {
		return entity.Post{}, result.Error
	}

	var storageObjects []entity.StorageObject

	result = database.DB.
		Table("storages").
		Select([]string{
			"id AS Id",
			"post_id AS PostId",
			"user_id AS UserId",
			"url AS Url",
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

func (repository PostRepository) FindAll(displayId string, search string, limit int, offset int) ([]entity.PostsItem, error) {
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
	sub := "SELECT id FROM storages WHERE post_id = posts.id LIMIT 1"

	result := database.DB.
		Table("posts").
		Select(strings.Join([]string{
			"posts.id AS PostId",
			"posts.user_id AS UserId",
			"posts.caption AS Caption",
			"posts.created_at AS CreatedAt",
			"storages.url AS Thumbnail",
			"users.display_id AS DisplayId",
			"users.avatar AS Avatar",
			"users.name AS Name",
		}, ",")).
		Joins("JOIN storages ON storages.id = (" + sub + ")").
		Joins("JOIN users ON users.id = posts.user_id")

	if displayId != "" {
		result = result.Where("posts.user_id = (SELECT id FROM users WHERE display_id = ?)", displayId)
	}

	if search != "" {
		result = result.Where("posts.caption LIKE ?", "%"+search+"%")
	}

	result = result.
		Order("posts.created_at DESC").
		Limit(limit).
		Offset(offset).
		Scan(&postsResult)

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

func (repository PostRepository) GetTotalCount(displayId string, search string) (int64, error) {
	var count int64
	result := database.DB.Table("posts")

	if displayId != "" {
		result = result.
			Where("posts.user_id = (SELECT id FROM users WHERE display_id = ?)", displayId)
	}

	if search != "" {
		result = result.Where("posts.caption LIKE ?", "%"+search+"%")
	}

	result = result.Count(&count)

	return count, result.Error
}
