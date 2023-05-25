package genreRouter

import (
	"Demo/internal/genre/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	r.POST("/genre", service.CreateGenre(db))
	r.GET("/genre/:id", service.GetGenreById(db))
	r.PATCH("/genre/:id", service.UpdateGenre(db))
	r.DELETE("genre/:id", service.DeleteGenre(db))
}
