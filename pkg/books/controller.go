package books

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type handler struct {
	DB       *gorm.DB
	validate *validator.Validate
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	v := validator.New()

	h := &handler{DB: db, validate: v}

	routes := router.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBooks)
	routes.GET("/:id", h.GetBook)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
