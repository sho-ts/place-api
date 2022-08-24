package entity

type UserDetail struct {
	Id        string `json:"userId"`
	DisplayId string `json:"displayId"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	Follow    int    `json:"follow"`
	Follower  int    `json:"follower"`
}

func NewUserDetail(
	id string,
	displayId string,
	name string,
	follow int,
	follower int,
) UserDetail {
	return UserDetail{
		Id:        id,
		DisplayId: displayId,
		Name:      name,
		Follow:    follow,
		Follower:  follower,
	}
}
