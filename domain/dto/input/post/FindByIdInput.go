package input

type FindByIdInput struct {
	PostId string
	UserId string
}

func NewFindByIdInput(
	postId string,
	userId string,
) FindByIdInput {
	return FindByIdInput{
		PostId: postId,
		UserId: userId,
	}
}
