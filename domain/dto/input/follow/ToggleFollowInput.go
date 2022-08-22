package input

type ToggleFollowInput struct {
	FollowUserId   string
	FollowerUserId string
}

func NewToggleFollowInput(followUserId string, followerUserId string) ToggleFollowInput {
	return ToggleFollowInput{
		FollowUserId:   followUserId,
		FollowerUserId: followerUserId,
	}
}
