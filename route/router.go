package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/controller"
	"github.com/sho-ts/place-api/route/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.GetCorsOption())

	pr := r.Group("/api/v1")
	ar := r.Group("/api/v1")
	ar.Use(middleware.GetAuthMiddleware().MiddlewareFunc())

  // Public Routes
	pr.GET("/hello", controller.GetHello)
	pr.POST("/user", controller.CreateUser)
	pr.GET("/user/:userId", controller.GetUser)
	pr.GET("/user/:userId/posts", controller.GetUserPosts)
	pr.GET("/user/duplicate/:userId", controller.CheckDuplicateUser)
	pr.GET("/post", controller.GetPosts)
	pr.GET("/post/:postId", controller.GetPost)
	pr.GET("/post/:postId/like/count", controller.GetLikeCount)
	pr.GET("/post/:postId/comment", controller.GetComments)

  // Auth Routes
	ar.GET("/user", controller.GetMe)
	ar.POST("/post", controller.CreatePost)
	ar.PUT("/post/like", controller.Like)
	ar.POST("/post/comment", controller.CreateComment)

	return r
}
