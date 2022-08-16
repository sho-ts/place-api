package input

type FindByIdInput struct {
	Id string
}

func NewFindByIdInput(
	id string,
) FindByIdInput {
	createUserInput := FindByIdInput{
		Id: id,
	}

	return createUserInput
}
