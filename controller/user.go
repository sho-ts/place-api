package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/service"
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
