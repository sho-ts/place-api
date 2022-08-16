package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/domain/dto/input/user"
	"github.com/sho-ts/place-api/usecase/user"
	"github.com/sho-ts/place-api/util"
)

type UserController struct {
	userCreateUseCase      usecase.ICreateUserUseCase
	FindByIdUseCase        usecase.IFindByIdUseCase
	FindByDisplayIdUseCase usecase.IFindByDisplayIdUseCase
}

func NewUserController(
	createUserUseCase usecase.ICreateUserUseCase,
	FindByIdUseCase usecase.IFindByIdUseCase,
	FindByDisplayIdUseCase usecase.IFindByDisplayIdUseCase,
) UserController {
	return UserController{
		userCreateUseCase:      createUserUseCase,
		FindByIdUseCase:        FindByIdUseCase,
		FindByDisplayIdUseCase: FindByDisplayIdUseCase,
	}
}

func (controller UserController) CreateUser(c *gin.Context) {
	var requestBody struct {
		UserId    string `json:"authId"`
		DisplayId string `json:"displayId"`
		Name      string `json:"name"`
	}
	c.ShouldBindJSON(&requestBody)

	FindByDisplayIdInput := input.NewFindByDisplayIdInput(requestBody.DisplayId)

	duplicate, _ := controller.FindByDisplayIdUseCase.Handle(FindByDisplayIdInput)

	if duplicate.Id != "" {
		c.JSON(500, gin.H{
			"message": "duplicate",
		})
    return
	}

	createUserInput := input.NewCreateUserInput(
		requestBody.UserId,
		requestBody.DisplayId,
		requestBody.Name,
	)

	user, err := controller.userCreateUseCase.Handle(createUserInput)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
    return
	}

	c.JSON(200, user)
}

func (controller UserController) GetMe(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	i := input.NewFindByIdInput(claims["sub"].(string))

	user, err := controller.FindByIdUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
    return
	}

	c.JSON(200, user)
}

func (controller UserController) GetUser(c *gin.Context) {
	i := input.NewFindByDisplayIdInput(c.Param("displayId"))

	user, err := controller.FindByDisplayIdUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
    return
	}

	c.JSON(200, user)
}
