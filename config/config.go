package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadConfig() *gorm.DB {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	envPath := filepath.Join(dir, "../../../../../../Desktop/API/config/.env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	return db
}

func RunFile(r *gin.Engine) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	envPath := filepath.Join(dir, "../../../../../../Desktop/API/config/.env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}
	dbPort := os.Getenv("DB_PORT_RUN")
	err = r.Run(dbPort)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
