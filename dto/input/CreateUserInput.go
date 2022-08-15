package input

import (
	"github.com/sho-ts/place-api/object"
)

type CreateUserInput struct {
	UserId    string
	DisplayId object.DisplayId
	Name      object.UserName
}

func NewCreateUserInput(
	userId string,
	displayId string,
	userName string,
) (CreateUserInput, error) {
	createUserInput := CreateUserInput{
		UserId: userId,
	}

	di := object.NewDisplayId(displayId)
	err := di.Valid()

	if err != nil {
		return createUserInput, err
	}

	un := object.NewUserName(userName)
	err = un.Valid()

	if err != nil {
		return createUserInput, err
	}

	createUserInput.Name = un
	createUserInput.DisplayId = di

	return createUserInput, nil
}
