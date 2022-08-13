package controller

import (
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

	var requestBody struct {
		Caption string `json:"caption"`
	}

	postId := util.GetULID()
	authId := claims["sub"].(string)

	msg := "投稿に失敗しました"
	_, err = service.CreatePost(postId, authId, requestBody.Caption)

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
