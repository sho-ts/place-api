package entity

type Post struct {
	Entity
	Id      string `gorm:"size:255;primary_key"`
	UserId  string
	Caption string
}
