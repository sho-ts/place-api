package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/service"
	"github.com/sho-ts/place-api/util"
)

func CreateUser(c *gin.Context) {
	var requestBody struct {
		AuthId    string `json:"authId"`
		Name      string `json:"name"`
		DisplayId string `json:"userId"`
	}

	c.ShouldBindJSON(&requestBody)

	user := service.CreateUser(requestBody.AuthId, requestBody.DisplayId, requestBody.Name)

	c.JSON(200, user)
}

func GetMe(c *gin.Context) {
	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	user, err := service.GetMe(claims["sub"].(string))

	if err != nil {
		c.JSON(404, gin.H{
			"message": "ユーザーが見つかりませんでした",
		})
		return
	}

	c.JSON(200, user)
}

func GetUser(c *gin.Context) {
	user, err := service.GetUser(c.Param("userId"))

	if err != nil {
		c.JSON(404, gin.H{
			"message": "ユーザーが見つかりませんでした",
		})
		return
	}

	c.JSON(200, user)
}