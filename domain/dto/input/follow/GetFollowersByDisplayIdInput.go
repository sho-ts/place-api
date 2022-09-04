package input

type GetFollowersByDisplayIdInput struct {
	DisplayId string
	UserId    string
	Limit     int
	Offset    int
}

func NewGetFollowersByDisplayIdInput(
	displayId string,
	userId string,
	limit int,
	offset int,
) GetFollowersByDisplayIdInput {
	return GetFollowersByDisplayIdInput{
		DisplayId: displayId,
		UserId:    userId,
		Limit:     limit,
		Offset:    offset,
	}
}
