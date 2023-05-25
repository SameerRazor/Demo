package libraryRouter

import (
	"Demo/internal/library/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	r.POST("/library", service.CreateLibrary(db))
	r.GET("/library/:id", service.GetLibrary(db))
	r.DELETE("library/:id", service.DeleteLibrary(db))
}