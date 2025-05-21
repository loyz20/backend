// internal/order/routes.go
package order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	usecase := NewUseCase(repo)
	handler := NewHandler(usecase)

	orderRoutes := rg.Group("/orders")
	{
		orderRoutes.POST("", handler.Create)
		orderRoutes.GET("", handler.GetAll)
		orderRoutes.GET(":id", handler.GetByID)
		orderRoutes.DELETE(":id", handler.Delete)
	}
}
