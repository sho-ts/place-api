package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/application/util"
	input "github.com/sho-ts/place-api/domain/dto/input/user"
	usecase "github.com/sho-ts/place-api/usecase/user"
)

type UserController struct {
	createUserUseCase      usecase.ICreateUserUseCase
	findByIdUseCase        usecase.IFindByIdUseCase
	findByDisplayIdUseCase usecase.IFindByDisplayIdUseCase
	changeProfileUseCase   usecase.IChangeProfileUseCase
}

func NewUserController(
	createUserUseCase usecase.ICreateUserUseCase,
	findByIdUseCase usecase.IFindByIdUseCase,
	findByDisplayIdUseCase usecase.IFindByDisplayIdUseCase,
	changeProfileUseCase usecase.IChangeProfileUseCase,
) UserController {
	return UserController{
		createUserUseCase:      createUserUseCase,
		findByIdUseCase:        findByIdUseCase,
		findByDisplayIdUseCase: findByDisplayIdUseCase,
		changeProfileUseCase:   changeProfileUseCase,
	}
}

func (controller UserController) CreateUser(c *gin.Context) {
	var requestBody struct {
		UserId    string `json:"authId"`
		DisplayId string `json:"displayId"`
		Name      string `json:"name"`
	}
	c.ShouldBindJSON(&requestBody)

	FindByDisplayIdInput := input.NewFindByDisplayIdInput(requestBody.DisplayId, "")

	duplicate, _ := controller.findByDisplayIdUseCase.Handle(FindByDisplayIdInput)

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

	user, err := controller.createUserUseCase.Handle(createUserInput)

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

	user, err := controller.findByIdUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, user)
}

func (controller UserController) GetUser(c *gin.Context) {
	i := input.NewFindByDisplayIdInput(c.Param("displayId"), c.Query("userId"))

	user, err := controller.findByDisplayIdUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, user)
}

func (controller UserController) ChangeProfile(c *gin.Context) {
  var fileName string
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	file, header, _ := c.Request.FormFile("attachmentFile")

  if header != nil {
    fileName = header.Filename
  }

	i := input.NewChangeProfileInput(
		claims["sub"].(string),
		c.Request.FormValue("displayId"),
		c.Request.FormValue("name"),
		file,
		fileName,
	)

	user, err := controller.changeProfileUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, user)
}
