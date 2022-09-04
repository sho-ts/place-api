package table

type Post struct {
	Id      string `gorm:"size:255;primary_key"`
	UserId  string
	Caption string
	Entity
}
