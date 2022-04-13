package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var (
// 	DB *gorm.DB
// )

func InitDB() *gorm.DB {
	// declare struct config & variable connectionString
	connectionString := os.Getenv("CONNECTION_DB")

	// var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return DB
}
