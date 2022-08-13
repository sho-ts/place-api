package entity

type Comment struct {
	Entity
	Id      string `gorm:"size:255;primary_key"`
	UserId  string
	PostId  string
	Content string
}
