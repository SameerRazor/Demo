package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadConfig() *gorm.DB {
	dsn := "root:J4C7ukpk@tcp(localhost:3306)/restapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}
