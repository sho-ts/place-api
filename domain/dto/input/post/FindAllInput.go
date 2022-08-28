package input

type FindAllInput struct {
	DisplayId string
  Limit int
  Offset int
}

func NewFindAllInput(
	displayId string,
  limit int,
  offset int,
) FindAllInput {
	return FindAllInput{
		DisplayId: displayId,
    Limit: limit,
    Offset: offset,
	}
}
