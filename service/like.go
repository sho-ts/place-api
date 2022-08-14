package service

import (
	"github.com/sho-ts/place-api/database"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
)

func AddLike(i input.HandleLikeInput) error {
	like := entity.Like{
		Id:     util.GetULID(),
		PostId: i.PostId,
		UserId: i.UserId,
	}

	result := database.DB.Create(&like)

	return result.Error
}

func RemoveLike(i input.HandleLikeInput) error {
	var like entity.Like

	// where
	pw := "post_id = ?"
	uw := "user_id = ?"

	result := database.DB.Where(pw, i.PostId).Where(uw, i.UserId).Delete(&like)

	return result.Error
}

func GetLikeCount(postId string) (output.CountOutput, error) {
	w := "post_id = ?"
	var count int64

	result := database.DB.Table("likes").Where(w, postId).Count(&count)

	o := output.CountOutput{
		Count: count,
	}

	return o, result.Error
}

func CheckDuplicateLike(i input.HandleLikeInput) (bool, error) {
	var count int64

	// where
	pw := "post_id = ?"
	uw := "user_id = ?"

	result := database.DB.Table("likes").Where(pw, i.PostId).Where(uw, i.UserId).Count(&count)

	if count > 0 {
		return true, result.Error
	} else {
		return false, result.Error
	}
}
