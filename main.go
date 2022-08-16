package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sho-ts/place-api/infrastructure/database"
	"github.com/sho-ts/place-api/route"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("環境変数の読み込みに失敗")
	}
}

func main() {
	loadEnv()
	database.Connect()
  database.Migrate()
	router := route.GetRouter()
	router.Run()
}
