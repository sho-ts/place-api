package input

type GetFollowsByDisplayIdInput struct {
	DisplayId string
	UserId    string
	Limit     int
	Offset    int
}

func NewGetFollowsByDisplayIdInput(
	displayId string,
	userId string,
	limit int,
	offset int,
) GetFollowsByDisplayIdInput {
	return GetFollowsByDisplayIdInput{
		DisplayId: displayId,
		UserId:    userId,
		Limit:     limit,
		Offset:    offset,
	}
}
