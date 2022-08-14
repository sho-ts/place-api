package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/constant"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/service"
	"github.com/sho-ts/place-api/util"
)

func CreateComment(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	var r struct {
		PostId  string `json:"postId"`
		Content string `json:"content"`
	}

	c.ShouldBindJSON(&r)

	i := input.CreateCommentInput{
		Id:      util.GetULID(),
		UserId:  claims["sub"].(string),
		PostId:  r.PostId,
		Content: r.Content,
	}

	comment, err := service.CreateComment(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_COMMENT_CREATE,
		})
		return
	}

	c.JSON(200, comment)
}

func GetComments(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)
	o, err := service.GetComments(c.Param("postId"), limit, offset)

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.FAILED_TO_GET_COMMENTS,
		})
		return
	}

	c.JSON(200, o)
}
