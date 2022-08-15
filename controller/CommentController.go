package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/constant"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
)

type ICommentService interface {
	CreateComment(i input.CreateCommentInput) (entity.Comment, error)
	GetComments(postId string, limit int, offset int) ([]output.GetCommentOutput, error)
}

type CommentController struct {
	commentService ICommentService
}

func NewCommentController(commentService ICommentService) CommentController {
	commentController := CommentController{
		commentService: commentService,
	}
	return commentController
}

func (cc CommentController) CreateComment(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	var r struct {
		PostId  string `json:"postId"`
		Content string `json:"content"`
	}

	c.ShouldBindJSON(&r)

	i, err := input.NewCreateCommentInput(util.GetULID(), claims["sub"].(string), r.PostId, r.Content)

  if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	comment, err := cc.commentService.CreateComment(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_COMMENT_CREATE,
		})
		return
	}

	c.JSON(200, comment)
}

func (cc CommentController) GetComments(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)
	o, err := cc.commentService.GetComments(c.Param("postId"), limit, offset)

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.FAILED_TO_GET_COMMENTS,
		})
		return
	}

	c.JSON(200, o)
}