package warehouse

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	usecase := NewUseCase(repo)
	handler := NewHandler(usecase)

	warehouseRoute := rg.Group("/warehouses")
	{
		warehouseRoute.GET("/", handler.GetAll)
		warehouseRoute.POST("/", handler.Create)
		warehouseRoute.GET("/:id", handler.GetByID)
		warehouseRoute.PUT("/:id", handler.Update)
		warehouseRoute.DELETE("/:id", handler.Delete)
	}
}
