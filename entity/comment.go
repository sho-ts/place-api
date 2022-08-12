package entity

type Comment struct {
	Entity
	UserId  string
	PostId  string
	Content string
}
