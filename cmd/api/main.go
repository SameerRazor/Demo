package main

import (
	"Demo/config"
	"Demo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.LoadConfig()

	r := gin.Default()

	config.AutoMigrateTable(db)

	router.SetupRouter(r, db)

	config.RunFile(r)
}
