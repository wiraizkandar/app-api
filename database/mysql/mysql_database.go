package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	dbInstance, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", os.Getenv("DB_NAME")))
	}

	return dbInstance
}
