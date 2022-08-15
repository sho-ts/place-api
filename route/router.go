package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/controller"
	"github.com/sho-ts/place-api/route/middleware"
	"github.com/sho-ts/place-api/service"
)

func GetRouter() *gin.Engine {
	userController := controller.NewUserController(service.NewUserService())
	postController := controller.NewPostController(service.NewPostService(), service.NewStorageService())
	likeController := controller.NewLikeController(service.NewLikeService())
	commentController := controller.NewCommentController(service.NewCommentService())

	r := gin.Default()
	r.Use(middleware.GetCorsOption())

	public := r.Group("/v1")

	public.POST("/users", userController.CreateUser)
	public.GET("/users/:userId", userController.GetUser)
	public.GET("/users/duplicate/:userId", userController.CheckDuplicateUser)
	public.GET("/users/:userId/posts", postController.GetUserPosts)
	public.GET("/posts", postController.GetPosts)
	public.GET("/posts/:postId", postController.GetPost)
	public.GET("/posts/:postId/like/count", likeController.GetLikeCount)
	public.GET("/posts/:postId/comment", commentController.GetComments)

	guard := r.Group("/v1")
	guard.Use(middleware.GetAuthMiddleware().MiddlewareFunc())

	guard.GET("/users", userController.GetMe)
	guard.POST("/posts", postController.CreatePost)
	guard.PUT("/posts/like", likeController.Like)
	guard.POST("/posts/comment", commentController.CreateComment)

	return r
}
