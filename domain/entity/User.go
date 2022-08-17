package entity

type User struct {
	Id        string `json:"-"`
	DisplayId string `json:"userId"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
}

func NewUser(
	id string,
	displayId string,
	name string,
) User {
	return User{
		Id:        id,
		DisplayId: displayId,
		Name:      name,
	}
}
