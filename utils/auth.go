package utils

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
)

func GetAuthResult(c *gin.Context) *jwtgo.Token {
	token, exists := c.Get("JWT_TOKEN")

	if !exists {
		panic("Error")
	}

	return token.(*jwtgo.Token)
}
