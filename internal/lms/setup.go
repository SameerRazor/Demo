package setup

import (
	"Demo/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeTable() (*gin.Engine, *gorm.DB) {
	router := gin.Default()

	db := config.LoadConfigTest()
	db.Exec("TRUNCATE TABLE authors;")
	db.Exec("TRUNCATE TABLE books;")
	db.Exec("TRUNCATE TABLE genres;")
	db.Exec("TRUNCATE TABLE libraries;")
	return router, db
}

func DeleteTables(db *gorm.DB){
	db.Exec("TRUNCATE TABLE authors;")
	db.Exec("TRUNCATE TABLE books;")
	db.Exec("TRUNCATE TABLE genres;")
	db.Exec("TRUNCATE TABLE libraries;")
}
