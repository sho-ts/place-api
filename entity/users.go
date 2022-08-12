package entity

type User struct {
	Entity
	AuthId    string
	DisplayId string
	Name      string
	Avatar    string
	Comment   Comment `gorm:"foreignkey:UserId"`
	Post      Post    `gorm:"foreignkey:UserId"`
	Storage   Storage `gorm:"foreignkey:UserId"`
}
