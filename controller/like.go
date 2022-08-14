package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/constant"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/service"
	"github.com/sho-ts/place-api/util"
)

func Like(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	var b struct {
		PostId string
	}

	c.ShouldBindJSON(&b)

	i := input.HandleLikeInput{
		PostId: b.PostId,
		UserId: claims["sub"].(string),
	}

	d, err := service.CheckDuplicateLike(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_LIKE,
		})
		return
	}

	if !d {
		err = service.AddLike(i)
	} else {
		err = service.RemoveLike(i)
	}

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_LIKE,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func GetLikeCount(c *gin.Context) {
	o, err := service.GetLikeCount(c.Param("postId"))

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.FAILED_TO_GET_LIKE,
		})
    return
	}

	c.JSON(200, o)
}
