package route

import (
	"github.com/gin-gonic/gin"
	app "github.com/sho-ts/place-api/application"
	"github.com/sho-ts/place-api/route/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.GetCorsOption())

	public := r.Group("/v1")

	public.POST("/users", app.UserController.CreateUser)
	public.GET("/users/:displayId", app.UserController.GetUser)

	public.GET("/posts", app.PostController.FindAll)
	public.GET("/posts/:postId", app.PostController.FindById)
	public.GET("/posts/:postId/comment", app.CommentController.FindAll)

	guard := r.Group("/v1")
	guard.Use(middleware.GetAuthMiddleware().MiddlewareFunc())

	guard.POST("/posts", app.PostController.CreatePost)
	guard.GET("/users", app.UserController.GetMe)
	guard.PUT("/posts/:postId/like", app.LikeController.ToggleLike)
	guard.POST("/posts/:postId/comment", app.CommentController.CreateComment)

	return r
}
