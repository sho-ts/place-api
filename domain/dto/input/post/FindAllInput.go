package input

type FindAllInput struct {
	UserId string
  Limit int
  Offset int
}

func NewFindAllInput(
	userId string,
  limit int,
  offset int,
) FindAllInput {
	return FindAllInput{
		UserId: userId,
    Limit: limit,
    Offset: offset,
	}
}
