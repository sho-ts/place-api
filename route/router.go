package route

import (
	"os"

	cognito "github.com/akhettar/gin-jwt-cognito"
	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/controller"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	mw, err := cognito.AuthJWTMiddleware(os.Getenv("AWS_COGNITO_ISS"), os.Getenv("AWS_COGNITO_USER_POOL_ID"), os.Getenv("AWS_COGNITO_REGION"))

	if err != nil {
		panic("router Error")
	}

	router.GET("/hello", controller.GetHello)
	router.POST("/user", mw.MiddlewareFunc(), controller.CreateUser)

	return router
}
