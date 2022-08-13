package entity

type User struct {
	Entity
	Id        string `json:"-" gorm:"size:255;primary_key"`
	DisplayId string `json:"userId"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
}
