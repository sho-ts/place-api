package input

type GetFollowersByDisplayIdInput struct {
	DisplayId string
	Limit     int
	Offset    int
}

func NewGetFollowersByDisplayIdInput(
	displayId string,
	limit int,
	offset int,
) GetFollowersByDisplayIdInput {
	return GetFollowersByDisplayIdInput{
		DisplayId: displayId,
		Limit:     limit,
		Offset:    offset,
	}
}
