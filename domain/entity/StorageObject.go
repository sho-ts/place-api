package entity

type StorageObject struct {
	Id     string `json:"id"`
	UserId string `json:"-"`
	PostId string `json:"-"`
	Url    string `json:"url"`
}

func NewStorageObject(
	id string,
	userId string,
	postId string,
	url string,
) StorageObject {
	return StorageObject{
		Id:     id,
		UserId: userId,
		PostId: postId,
		Url:    url,
	}
}
