package input

type GetFollowsByDisplayIdInput struct {
	DisplayId string
	Limit     int
	Offset    int
}

func NewGetFollowsByDisplayIdInput(
	displayId string,
	limit int,
	offset int,
) GetFollowsByDisplayIdInput {
	return GetFollowsByDisplayIdInput{
		DisplayId: displayId,
		Limit:     limit,
		Offset:    offset,
	}
}
