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

	pr := r.Group("/v1")
	ar := r.Group("/v1")
	ar.Use(middleware.GetAuthMiddleware().MiddlewareFunc())

	pr.POST("/user", userController.CreateUser)
	pr.GET("/user/:userId", userController.GetUser)
	pr.GET("/user/duplicate/:userId", userController.CheckDuplicateUser)
	ar.GET("/user", userController.GetMe)
	pr.GET("/user/:userId/posts", postController.GetUserPosts)
  
  pr.GET("/post", postController.GetPosts)
	pr.GET("/post/:postId", postController.GetPost)
	ar.POST("/post", postController.CreatePost)

	pr.GET("/post/:postId/like/count", likeController.GetLikeCount)
	ar.PUT("/post/like", likeController.Like)

	pr.GET("/post/:postId/comment", commentController.GetComments)
	ar.POST("/post/comment", commentController.CreateComment)

  return r
}
