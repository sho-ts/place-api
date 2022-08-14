package entity

type Storage struct {
	Entity
	Id     string `gorm:"size:255;primary_key" json:"id"`
	UserId string `json:"-"`
	PostId string `json:"-"`
	Url    string `json:"url"`
}
