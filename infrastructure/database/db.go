package database

import (
  "github.com/sho-ts/place-api/infrastructure/database/table"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
  "strings"
)

var DB *gorm.DB

func Connect() {
	var err error

  dsn := strings.Join([]string{
    os.Getenv("MYSQL_USER") + ":",
    os.Getenv("MYSQL_PASSWORD"),
    "@tcp(" + os.Getenv("MYSQL_HOST") + ")/",
    os.Getenv("MYSQL_DBNAME"),
    "?parseTime=true",
  },"")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func Migrate() {
	DB.AutoMigrate(
		&table.User{},
		&table.Post{},
		&table.Comment{},
		&table.Like{},
		&table.Storage{},
    &table.Follow{},
	)
}
