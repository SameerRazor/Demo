package bookRouter

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Demo/internal/book/service"
)

func SetupRouter(r *gin.Engine, db *gorm.DB){
	r.GET("/books", service.GetBooks(db))
	r.GET("/books/:params", service.GetBookParams(db, "params"))
	r.POST("/books", service.CreateBooks(db))
	r.PATCH("/books/:id", service.UpdateBooks(db))
	r.DELETE("/books/:id", service.DeleteBook(db))
}
