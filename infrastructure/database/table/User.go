package table

type User struct {
  Id        string `gorm:"size:255;primary_key;unique" json:"-"`
	DisplayId string `json:"userId"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Entity
}
