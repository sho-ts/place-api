package input

type FindByDisplayIdInput struct {
	DisplayId string
}

func NewFindByDisplayIdInput(
	displayId string,
) FindByDisplayIdInput {
	createUserInput := FindByDisplayIdInput{
		DisplayId: displayId,
	}

	return createUserInput
}
