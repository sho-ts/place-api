package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/constant"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/util"
)

type ILikeService interface {
	AddLike(i input.HandleLikeInput) error
	RemoveLike(i input.HandleLikeInput) error
	GetLikeCount(postId string) (output.CountOutput, error)
	CheckDuplicateLike(i input.HandleLikeInput) (bool, error)
}

type LikeController struct {
	likeService ILikeService
}

func NewLikeController(likeService ILikeService) LikeController {
	likeController := LikeController{
		likeService: likeService,
	}
	return likeController
}

func (lc LikeController) Like(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	var b struct {
		PostId string
	}

	c.ShouldBindJSON(&b)

  i := input.NewHandleLikeInput(
    b.PostId,
    claims["sub"].(string),
  )

	d, err := lc.likeService.CheckDuplicateLike(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_LIKE,
		})
		return
	}

	if !d {
		err = lc.likeService.AddLike(i)
	} else {
		err = lc.likeService.RemoveLike(i)
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

func (lc LikeController) GetLikeCount(c *gin.Context) {
	o, err := lc.likeService.GetLikeCount(c.Param("postId"))

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.FAILED_TO_GET_LIKE,
		})
		return
	}

	c.JSON(200, o)
}