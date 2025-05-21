package product

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	usecase := NewUseCase(repo)
	handler := NewHandler(usecase)

	productRoute := rg.Group("/products")
	{
		productRoute.GET("/", handler.GetAll)
		productRoute.POST("/", handler.Create)
		productRoute.GET("/:id", handler.GetByID)
		productRoute.PUT("/:id", handler.Update)
		productRoute.DELETE("/:id", handler.Delete)
	}
}
