package controller

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/constant"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/util"
)

type IPostService interface {
	GetPosts(search string, limit int, offset int) (output.GetPostsOutput, error)
	GetUserPosts(userId string, limit int, offset int) (output.GetPostsOutput, error)
}

type IStorageService interface {
	UploadToS3Bucket(file multipart.File, name string) (string, error)
}

type PostController struct {
	postService    IPostService
	storageService IStorageService
}

func NewPostController(
	postService IPostService,
	storageService IStorageService,
) PostController {
	postController := PostController{
		postService:    postService,
		storageService: storageService,
	}
	return postController
}

func (pc PostController) GetPosts(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)

	o, err := pc.postService.GetPosts(c.Query("s"), limit, offset)

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.NOT_FOUND_POST,
		})
		return
	}

	c.JSON(200, o)
}

func (pc PostController) GetUserPosts(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)

	o, err := pc.postService.GetUserPosts(c.Param("userId"), limit, offset)

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.NOT_FOUND_POST,
		})
		return
	}

	c.JSON(200, o)
}
