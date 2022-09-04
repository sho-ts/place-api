package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/application/util"
	"github.com/sho-ts/place-api/domain/dto/input/follow"
	"github.com/sho-ts/place-api/usecase/follow"
)

type FollowController struct {
	toggleFollowUseCase            usecase.IToggleFollowUseCase
	getFollowsByDisplayIdUseCase   usecase.IGetFollowsByDisplayIdUseCase
	getFollowersByDisplayIdUseCase usecase.IGetFollowersByDisplayIdUseCase
}

func NewFollowController(
	toggleFollowUseCase usecase.IToggleFollowUseCase,
	getFollowsByDisplayIdUseCase usecase.IGetFollowsByDisplayIdUseCase,
	getFollowersByDisplayIdUseCase usecase.IGetFollowersByDisplayIdUseCase,
) FollowController {
	return FollowController{
		toggleFollowUseCase:            toggleFollowUseCase,
		getFollowsByDisplayIdUseCase:   getFollowsByDisplayIdUseCase,
		getFollowersByDisplayIdUseCase: getFollowersByDisplayIdUseCase,
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

func (controller FollowController) GetFollowsByDisplayId(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)

	i := input.NewGetFollowsByDisplayIdInput(
		c.Param("displayId"),
		c.Query("userId"),
		limit,
		offset,
	)

	o, err := controller.getFollowsByDisplayIdUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, o)
}

func (controller FollowController) GetFollowersByDisplayId(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)

	i := input.NewGetFollowersByDisplayIdInput(
		c.Param("displayId"),
		c.Query("userId"),
		limit,
		offset,
	)

	o, err := controller.getFollowersByDisplayIdUseCase.Handle(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error",
		})
		return
	}

	c.JSON(200, o)
}
