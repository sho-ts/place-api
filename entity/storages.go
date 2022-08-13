package entity

type Storage struct {
	Entity
	Id     string `gorm:"size:255;primary_key"`
	UserId string `json:"-"`
	PostId string `json:"-"`
	Url    string
}
