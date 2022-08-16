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
  FindByIdUseCase usecase.IFindByIdUseCase
}

func NewPostController(
	createPostUseCase usecase.ICreatePostUseCase,
  findByIdUseCase usecase.IFindByIdUseCase,
) PostController {
	return PostController{
		CreatePostUseCase: createPostUseCase,
    FindByIdUseCase: findByIdUseCase,
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
    return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (controller PostController) FindById(c *gin.Context) {
  i := input.NewFindByIdInput(
    c.Param("postId"),
    c.Query("userId"),
  )

  post, err := controller.FindByIdUseCase.Handle(i)

  if err != nil {
    c.JSON(500, gin.H{
			"message": "Error",
		})
    return
  }

  c.JSON(200, post)
}