package entity

type UserDetail struct {
	Id           string `json:"userId"`
	DisplayId    string `json:"displayId"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	FollowStatus int    `json:"followStatus"`
}

func NewUserDetail(
	id string,
	displayId string,
	name string,
) UserDetail {
	return UserDetail{
		Id:        id,
		DisplayId: displayId,
		Name:      name,
	}
}
