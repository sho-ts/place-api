package route

import (
	"github.com/gin-gonic/gin"
	app "github.com/sho-ts/place-api/application"
	"github.com/sho-ts/place-api/controller"
	"github.com/sho-ts/place-api/route/middleware"
	"github.com/sho-ts/place-api/service"
)

func GetRouter() *gin.Engine {
	postController := controller.NewPostController(service.NewPostService(), service.NewStorageService())
	likeController := controller.NewLikeController(service.NewLikeService())
	commentController := controller.NewCommentController(service.NewCommentService())

	r := gin.Default()
	r.Use(middleware.GetCorsOption())

	public := r.Group("/v1")

	public.POST("/users", app.UserController.CreateUser)
	public.GET("/users/:displayId", app.UserController.GetUser)
  
	public.GET("/users/:displayId/posts", postController.GetUserPosts)
	public.GET("/posts", postController.GetPosts)
	public.GET("/posts/:postId", app.PostController.FindById)
	public.GET("/posts/:postId/like/count", likeController.GetLikeCount)
	public.GET("/posts/:postId/comment", commentController.GetComments)
  
	guard := r.Group("/v1")
	guard.Use(middleware.GetAuthMiddleware().MiddlewareFunc())
  
	guard.POST("/posts", app.PostController.CreatePost)
	guard.GET("/users", app.UserController.GetMe)
	guard.PUT("/posts/like", likeController.Like)
	guard.POST("/posts/comment", commentController.CreateComment)

	return r
}
