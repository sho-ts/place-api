package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func GetCorsOption() gin.HandlerFunc {
	cors := cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("CLIENT_WEB_URL"),
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authentication",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	})

	return cors
}
