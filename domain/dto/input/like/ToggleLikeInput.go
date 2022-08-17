package input

type ToggleLikeInput struct {
	PostId string
	UserId string
}

func NewToggleLikeInput(
	postId string,
	userId string,
) ToggleLikeInput {
	return ToggleLikeInput{
		PostId: postId,
		UserId: userId,
	}
}
