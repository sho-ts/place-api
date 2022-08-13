package entity

type Storage struct {
	Entity
	Id     string `gorm:"size:255;primary_key"`
	UserId string
	PostId string
	Url    string
}
