package input

type CreateUserInput struct {
	UserId    string
	DisplayId string
	Name      string
}

func NewCreateUserInput(
	userId string,
	displayId string,
	userName string,
) CreateUserInput {
	createUserInput := CreateUserInput{
		UserId:    userId,
		DisplayId: displayId,
		Name:      userName,
	}

	return createUserInput
}
