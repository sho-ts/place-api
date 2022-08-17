package input

type FindByDisplayIdInput struct {
	DisplayId string
}

func NewFindByDisplayIdInput(
	displayId string,
) FindByDisplayIdInput {
	return FindByDisplayIdInput{
		DisplayId: displayId,
	}
}
