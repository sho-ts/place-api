package entity

type Post struct {
	Entity
	Id      string `gorm:"size:255;primary_key"`
	UserId  string
	Name    string
	Avatar  string
	Caption string
}
