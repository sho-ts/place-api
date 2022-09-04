package middleware

import (
	cognito "github.com/akhettar/gin-jwt-cognito"
	"os"
)

func GetAuthMiddleware() *cognito.AuthMiddleware {
	mw, err := cognito.AuthJWTMiddleware(os.Getenv("AWS_COGNITO_ISS"), os.Getenv("AWS_COGNITO_USER_POOL_ID"), os.Getenv("AWS_COGNITO_REGION"))

	if err != nil {
		panic("router Error")
	}

	return mw
}
