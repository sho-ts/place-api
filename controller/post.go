package controller

import (
	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/constant"
	"github.com/sho-ts/place-api/dto/input"
	"github.com/sho-ts/place-api/service"
	"github.com/sho-ts/place-api/util"
)

func CreatePost(c *gin.Context) {
	file, header, _ := c.Request.FormFile("attachmentFile")
	path, err := service.UploadToS3Bucket(file, header.Filename)

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

	_, err = service.CreatePost(i)

	if err != nil {
		c.JSON(500, gin.H{
			"message": constant.FAILED_TO_POST_CREATE,
		})
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

func GetPost(c *gin.Context) {
	o, err := service.GetPost(c.Param("postId"))

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.NOT_FOUND_POST,
		})
	}

	c.JSON(200, o)
}

func GetPosts(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)

	o, err := service.GetPosts(c.Query("s"), limit, offset)

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.NOT_FOUND_POST,
		})
		return
	}

	c.JSON(200, o)
}

func GetUserPosts(c *gin.Context) {
	limit, offset := util.GetLimitAndOffset(c)

	o, err := service.GetUserPosts(c.Param("userId"), limit, offset)

	if err != nil {
		c.JSON(404, gin.H{
			"message": constant.NOT_FOUND_POST,
		})
		return
	}

	c.JSON(200, o)
}
