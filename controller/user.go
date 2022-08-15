package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/constant"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
)

type IUserService interface {
	CreateUser(i input.CreateUserInput) (entity.User, error)
  GetMe(authId string) (entity.User, error)
  GetUser(userId string) (entity.User, error)
}

type UserController struct {
	userService IUserService
}

func NewUserController(userService IUserService) UserController {
	userController := UserController{
		userService: userService,
	}
	return userController
}

func (uc UserController) CreateUser(c *gin.Context) {
	var i input.CreateUserInput
	c.ShouldBindJSON(&i)

	user, err := uc.userService.CreateUser(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_USER_CREATE,
		})
		return
	}

	c.JSON(200, user)
}

func (uc UserController) GetMe(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	user, err := uc.userService.GetMe(claims["sub"].(string))

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.NOT_FOUND_USER,
		})
		return
	}

	c.JSON(200, user)
}

func (uc UserController) GetUser(c *gin.Context) {
	user, err := uc.userService.GetUser(c.Param("userId"))

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.NOT_FOUND_USER,
		})
		return
	}

	c.JSON(200, user)
}

/* ユーザーの重複を確認する */
func (uc UserController) CheckDuplicateUser(c *gin.Context) {
	_, err := uc.userService.GetUser(c.Param("userId"))

	if err != nil {
		c.JSON(200, gin.H{
			"message": constant.NOT_FOUND_USER,
		})
		return
	}

	c.JSON(500, gin.H{
		"message": constant.DUPLICATE_USER,
	})
}
