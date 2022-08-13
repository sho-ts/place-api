package entity

type User struct {
	Entity
	Id        string `gorm:"size:255;primary_key"`
	AuthId    string
	DisplayId string
	Name      string
	Avatar    string
}
