package table

type Comment struct {
	Id      string `gorm:"size:255;primary_key"`
	UserId  string
	PostId  string
	Content string
	Entity
}
