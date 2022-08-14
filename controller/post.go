package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	jwtgo "github.com/golang-jwt/jwt"
	"github.com/sho-ts/place-api/service"
	"github.com/sho-ts/place-api/util"
)

func CreatePost(c *gin.Context) {
	file, header, _ := c.Request.FormFile("attachmentFile")
	path, err := service.UploadToS3Bucket(file, header.Filename)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "ファイルのアップロードに失敗しました",
		})
	}

	token := util.GetAuthResult(c)
	claims := token.Claims.(jwtgo.MapClaims)

	caption := c.Request.FormValue("caption")
	postId := util.GetULID()
	authId := claims["sub"].(string)

	msg := "投稿に失敗しました"
	_, err = service.CreatePost(postId, authId, caption)

	if err != nil {
		c.JSON(500, gin.H{
			"message": msg,
		})
	}

	_, err = service.CreateStorage(postId, authId, path)

	if err != nil {
		c.JSON(500, gin.H{
			"message": msg,
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
			"message": "投稿が見つかりませんでした",
		})
	}

	c.JSON(200, o)
}

func GetUserPosts(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil || limit > 30 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		offset = 0
	}

	o, err := service.GetUserPosts(c.Param("userId"), limit, offset)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "投稿が見つかりませんでした",
		})
		return
	}

	c.JSON(200, o)
}
