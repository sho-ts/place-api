package input

type HandleLikeInput struct {
	PostId string
	UserId string
}

func NewHandleLikeInput(
	postId string,
	userId string,
) HandleLikeInput {
	return HandleLikeInput{
		PostId: postId,
		UserId: userId,
	}
}
