package database

import (
	"github.com/sho-ts/place-api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("MYSQL_HOST") + ")/" + os.Getenv("MYSQL_DBNAME") + "?parseTime=true"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func Migrate() {
	DB.AutoMigrate(
		&entity.User{},
		&entity.Post{},
		&entity.Comment{},
		&entity.Like{},
		&entity.Storage{},
	)
}
