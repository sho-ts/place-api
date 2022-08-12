package entity

type Post struct {
	Entity
	UserId  string
	Name    string
	Avatar  string
	Storage Storage `gorm:"foreignkey:PostId"`
}
