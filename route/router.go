package route

import (
	cognito "github.com/akhettar/gin-jwt-cognito"
	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/controller"
	"github.com/sho-ts/place-api/route/middleware"
	"os"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.GetCorsOption())

	mw, err := cognito.AuthJWTMiddleware(os.Getenv("AWS_COGNITO_ISS"), os.Getenv("AWS_COGNITO_USER_POOL_ID"), os.Getenv("AWS_COGNITO_REGION"))

	if err != nil {
		panic("router Error")
	}

	router.GET("/api/v1/hello", controller.GetHello)

	router.GET("/api/v1/user", mw.MiddlewareFunc(), controller.GetMe)
	router.POST("/api/v1/user", controller.CreateUser)
	router.GET("/api/v1/user/:userId", controller.GetUser)
	router.GET("/api/v1/user/duplicate/:userId", controller.CheckDuplicateUser)

	router.GET("/api/v1/post/:postId", controller.GetPost)
	router.POST("/api/v1/post", mw.MiddlewareFunc(), controller.CreatePost)

	return router
}
