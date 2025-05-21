package customer

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	usecase := NewUseCase(repo)
	handler := NewHandler(usecase)

	customerRoute := rg.Group("/customers")
	{
		customerRoute.GET("/", handler.GetAll)
		customerRoute.GET("/:id", handler.GetByID)
		customerRoute.POST("/", handler.Create)
		customerRoute.PUT("/:id", handler.Update)
		customerRoute.DELETE("/:id", handler.Delete)
	}
}
