package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/application/util"
	"github.com/sho-ts/place-api/domain/dto/input/follow"
	"github.com/sho-ts/place-api/usecase/follow"
)

type FollowController struct {
	toggleFollowUseCase usecase.IToggleFollowUseCase
}

func NewFollowController(
	toggleFollowUseCase usecase.IToggleFollowUseCase,
) FollowController {
	return FollowController{
		toggleFollowUseCase: toggleFollowUseCase,
	}
}

func (controller FollowController) ToggleFollow(c *gin.Context) {
  var requestBody struct {
    FollowUserId string `json:"followUserId"`
  }
	c.ShouldBindJSON(&requestBody)

	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

  i := input.NewToggleFollowInput(
    requestBody.FollowUserId,
    claims["sub"].(string),
  )

  err := controller.toggleFollowUseCase.Handle(i)

  if err != nil {
    c.JSON(500, gin.H{
      "message": "Error",
    })
  }

  c.JSON(200, gin.H{
    "message": "success",
  })
}
