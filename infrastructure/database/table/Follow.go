package table

import (
	"time"
)

type Follow struct {
	Id             string `gorm:"size:255;primary_key"`
	FollowUserId   string
	FollowerUserId string
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
