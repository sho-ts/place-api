package entity

import (
	"time"
)

type Like struct {
  Id        string    `gorm:"size:255;primary_key"`
	UserId    string
	PostId    string
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
