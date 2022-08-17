package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/application/util"
	input "github.com/sho-ts/place-api/domain/dto/input/comment"
	usecase "github.com/sho-ts/place-api/usecase/comment"
)

type CommentController struct {
	CreateCommentUseCase usecase.ICreateCommentUseCase
	FindAllUseCase       usecase.IFindAllUseCase
}

func NewCommentController(
	createCommentUseCase usecase.ICreateCommentUseCase,
	findAllUseCase usecase.IFindAllUseCase,
) CommentController {
	return CommentController{
		CreateCommentUseCase: createCommentUseCase,
		FindAllUseCase:       findAllUseCase,
	}
}

func (controller CommentController) CreateComment(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	var r struct {
		Content string `json:"content"`
	}

	c.ShouldBindJSON(&r)

	i := input.NewCreateCommentInput(
		util.GetULID(),
		claims["sub"].(string),
		c.Param("postId"),
		r.Content,
	)

	err := controller.CreateCommentUseCase.Handle(i)

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

func (controller CommentController) FindAll(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)

	i := input.NewFindAllInput(
		c.Param("postId"),
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
