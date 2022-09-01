package input

type FindByDisplayIdInput struct {
	DisplayId string
	UserId    string
}

func NewFindByDisplayIdInput(
	displayId string,
	userId string,
) FindByDisplayIdInput {
	return FindByDisplayIdInput{
		DisplayId: displayId,
		UserId:    userId,
	}
}
