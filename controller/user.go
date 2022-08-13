package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/service"
	"github.com/sho-ts/place-api/util"
)

func CreateUser(c *gin.Context) {
	token := util.GetAuthResult(c)
	clims := token.Claims.(jwtgo.MapClaims)

	var requestBody struct {
		Name      string `json:"name"`
		DisplayId string `json:"userId"`
	}

	c.ShouldBindJSON(&requestBody)

	user := service.CreateUser(clims["sub"].(string), requestBody.DisplayId, requestBody.Name)

	c.JSON(200, user)
}