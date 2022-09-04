package entity

type Follow struct {
	Id           string `json:"userId"`
	DisplayId    string `json:"displayId"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	FollowStatus int    `json:"followStatus"`
}
