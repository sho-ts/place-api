package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/domain/dto/input/comment"
	"github.com/sho-ts/place-api/usecase/comment"
	"github.com/sho-ts/place-api/util"
)

type CommentController struct {
	CreateCommentUseCase usecase.ICreateCommentUseCase
}

func NewCommentController(
	createCommentUseCase usecase.ICreateCommentUseCase,
) CommentController {
	return CommentController{
		CreateCommentUseCase: createCommentUseCase,
	}
}

func (controller CommentController) CreateComment(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	var r struct {
		PostId  string `json:"postId"`
		Content string `json:"content"`
	}

	c.ShouldBindJSON(&r)

	i := input.NewCreateCommentInput(
		util.GetULID(),
		claims["sub"].(string),
		r.PostId,
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
