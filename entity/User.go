package entity

type User struct {
  Id        string `json:"-" gorm:"size:255;primary_key" gorm:"unique"`
	DisplayId string `json:"userId"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Entity
}
