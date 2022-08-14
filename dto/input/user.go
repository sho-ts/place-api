package input

type CreateUserInput struct {
	UserId    string `json:"authId"`
	DisplayId string `json:"userId"`
	Name      string `json:"name"`
}
