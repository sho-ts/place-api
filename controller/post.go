package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/constant"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/dto/output"
	"github.com/sho-ts/place-api/entity"
	"github.com/sho-ts/place-api/util"
	"mime/multipart"
)

type IPostService interface {
	CreatePost(i input.CreatePostInput) (entity.Post, error)
	GetPost(postId string, userId string) (output.GetPostOutput, error)
	GetPosts(search string, limit int, offset int) ([]output.GetPostsOutput, error)
	GetUserPosts(userId string, limit int, offset int) ([]output.GetPostsOutput, error)
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

func (pc PostController) CreatePost(c *gin.Context) {
	file, header, _ := c.Request.FormFile("attachmentFile")
	path, err := pc.storageService.UploadToS3Bucket(file, header.Filename)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_S3_UPLOAD,
		})
	}

	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	i := input.CreatePostInput{
		PostId:  util.GetULID(),
		UserId:  claims["sub"].(string),
		Caption: c.Request.FormValue("caption"),
		Urls:    []string{path},
	}

	_, err = pc.postService.CreatePost(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_POST_CREATE,
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func (pc PostController) GetPost(c *gin.Context) {
	o, err := pc.postService.GetPost(c.Param("postId"), c.Query("userId"))

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.NOT_FOUND_POST,
		})
	}

	c.JSON(200, o)
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
