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

	user := service.GetMe(claims["sub"].(string))

	c.JSON(200, user)
}
