package route

import (
	cognito "github.com/akhettar/gin-jwt-cognito"
	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/controller"
	"github.com/sho-ts/place-api/route/middleware"
	"os"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.GetCorsOption())

	mw, err := cognito.AuthJWTMiddleware(os.Getenv("AWS_COGNITO_ISS"), os.Getenv("AWS_COGNITO_USER_POOL_ID"), os.Getenv("AWS_COGNITO_REGION"))

	if err != nil {
		panic("router Error")
	}

	r.GET("/api/v1/hello", controller.GetHello)

	r.GET("/api/v1/user", mw.MiddlewareFunc(), controller.GetMe)
	r.POST("/api/v1/user", controller.CreateUser)
	r.GET("/api/v1/user/:userId", controller.GetUser)
	r.GET("/api/v1/user/:userId/posts", controller.GetUserPosts)
	r.GET("/api/v1/user/duplicate/:userId", controller.CheckDuplicateUser)

	r.GET("/api/v1/post", controller.GetPosts)
	r.POST("/api/v1/post", mw.MiddlewareFunc(), controller.CreatePost)
	r.GET("/api/v1/post/:postId", controller.GetPost)
	r.GET("/api/v1/post/:postId/like/count", controller.GetLikeCount)
	r.PUT("/api/v1/post/like", mw.MiddlewareFunc(), controller.Like)

	r.GET("/api/v1/post/:postId/comment", controller.GetComments)
	r.POST("/api/v1/post/comment", mw.MiddlewareFunc(), controller.CreateComment)

	return r
}
