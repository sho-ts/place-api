package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/application/util"
	input "github.com/sho-ts/place-api/domain/dto/input/like"
	usecase "github.com/sho-ts/place-api/usecase/like"
)

type LikeController struct {
	ToggleLikeUseCase usecase.IToggleLikeUseCase
}

func NewLikeController(
	toggleLikeUseCase usecase.IToggleLikeUseCase,
) LikeController {
	return LikeController{
		ToggleLikeUseCase: toggleLikeUseCase,
	}
}

func (controller LikeController) ToggleLike(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	i := input.NewToggleLikeInput(
		c.Param("postId"),
		claims["sub"].(string),
	)

	err := controller.ToggleLikeUseCase.Handle(i)

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
