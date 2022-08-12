package service

import (
	"github.com/gin-gonic/gin"
)

func GetHello() gin.H {
	return gin.H{
		"message": "hello",
	}
}