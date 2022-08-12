package entity

import (
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	Id        string         `gorm:"size:255;primary_key"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
