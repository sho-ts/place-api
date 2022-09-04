package repository

import (
	"github.com/sho-ts/place-api/domain/entity"
)

type IFollowRepository interface {
	Store(followUserId string, followerUserId string) error
	Remove(followUserId string, followerUserId string) error
	CheckDuplicate(postId string, userId string) (bool, error)
	GetFollowsByDisplayId(displayId string, userId string, limit int, offset int) ([]entity.Follow, int64, error)
	GetFollowersByDisplayId(displayId string, userId string, limit int, offset int) ([]entity.Follow, int64, error)
}
