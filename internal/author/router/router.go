package authorRouter

import (
	"Demo/internal/author/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	r.POST("/author", service.CreateAuthor(db))
	r.GET("/author", service.GetAuthor(db))
	r.GET("/author/:params", service.GetAuthorParams(db, "params"))
	r.PATCH("/author/:id", service.UpdateAuthor(db))
	r.DELETE("/author/:id", service.DeleteAuthor(db))
}