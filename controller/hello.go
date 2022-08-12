package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/service"
)

func GetHello(c *gin.Context) {
	c.JSON(200, service.GetHello())
}