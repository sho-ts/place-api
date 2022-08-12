package entity

type User struct {
	Entity
	AuthId    string
	DisplayId string
	Name      string
	Avatar    string
	Post      Post    `gorm:"foreignkey:UserId"`
	Storage   Storage `gorm:"foreignkey:UserId"`
}