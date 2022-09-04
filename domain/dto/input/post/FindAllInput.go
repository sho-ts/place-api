package input

type FindAllInput struct {
	DisplayId string
	Search    string
	Limit     int
	Offset    int
}

func NewFindAllInput(
	displayId string,
	search string,
	limit int,
	offset int,
) FindAllInput {
	return FindAllInput{
		DisplayId: displayId,
		Search:    search,
		Limit:     limit,
		Offset:    offset,
	}
}
