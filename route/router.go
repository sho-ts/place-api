package route

import (
	"github.com/gin-gonic/gin"
	"github.com/sho-ts/place-api/controller"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/hello", controller.GetHello)

	return router
}
