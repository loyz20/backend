// internal/supplier/routes.go
package supplier

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	usecase := NewUseCase(repo)
	handler := NewHandler(usecase)

	supplierRoute := rg.Group("/suppliers")
	{
		supplierRoute.GET("/", handler.GetAll)
		supplierRoute.GET("/:id", handler.GetByID)
		supplierRoute.POST("/", handler.Create)
		supplierRoute.PUT("/:id", handler.Update)
		supplierRoute.DELETE("/:id", handler.Delete)
	}
}
