package input

type FindByIdInput struct {
	Id string
}

func NewFindByIdInput(
	id string,
) FindByIdInput {
	return FindByIdInput{
		Id: id,
	}
}
