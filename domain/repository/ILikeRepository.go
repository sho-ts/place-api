package repository

type ILikeRepository interface {
	Store(postId string, userId string) error
	Remove(postId string, userId string) error
	CheckDuplicate(postId string, userId string) (bool, error)
}
