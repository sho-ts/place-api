package repository

type IFollowRepository interface {
	Store(followUserId string, followerUserId string) error
	Remove(followUserId string, followerUserId string) error
}
