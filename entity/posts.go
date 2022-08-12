package entity

type Post struct {
	Entity
	UserId  string
	Name    string
	Avatar  string
	Caption string
	Storage Storage `gorm:"foreignkey:PostId"`
	Comment Comment `gorm:"foreignkey:PostId"`
}
