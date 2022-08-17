package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/application/util"
	input "github.com/sho-ts/place-api/domain/dto/input/post"
	usecase "github.com/sho-ts/place-api/usecase/post"
)

type PostController struct {
	CreatePostUseCase usecase.ICreatePostUseCase
	FindByIdUseCase   usecase.IFindByIdUseCase
	FindAllUseCase   usecase.IFindAllUseCase
}

func NewPostController(
	createPostUseCase usecase.ICreatePostUseCase,
	findByIdUseCase usecase.IFindByIdUseCase,
	findAllUseCase usecase.IFindAllUseCase,
) PostController {
	return PostController{
		CreatePostUseCase: createPostUseCase,
		FindByIdUseCase:   findByIdUseCase,
		FindAllUseCase:   findAllUseCase,
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

func (controller PostController) FindAll(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)

	i := input.NewFindAllInput(
		c.Query("userId"),
		limit,
		offset,
	)

	o, err := controller.FindAllUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, o)
}
