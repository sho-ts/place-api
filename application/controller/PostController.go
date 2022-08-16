package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/domain/dto/input/post"
	"github.com/sho-ts/place-api/usecase/post"
	"github.com/sho-ts/place-api/util"
)

type PostController struct {
	CreatePostUseCase usecase.ICreatePostUseCase
}

func NewPostController(
	createPostUseCase usecase.ICreatePostUseCase,
) PostController {
	return PostController{
		CreatePostUseCase: createPostUseCase,
	}
}

func (controller PostController) CreatePost(c *gin.Context) {
	file, header, _ := c.Request.FormFile("attachmentFile")
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	i := input.NewCreatePostInput(
		util.GetULID(),
		claims["sub"].(string),
		c.Request.FormValue("caption"),
		file,
		header.Filename,
	)

	err := controller.CreatePostUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}
