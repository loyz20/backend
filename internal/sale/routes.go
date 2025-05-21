package sale

import (
	"backend/internal/product"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	productRepo := product.NewRepository(db)
	productUC := product.NewUseCase(productRepo)
	usecase := NewUseCase(repo, productUC)
	handler := NewHandler(usecase)

	saleRoutes := rg.Group("/sales")
	{
		saleRoutes.POST("", handler.Create)
		saleRoutes.GET("", handler.GetAll)
		saleRoutes.GET(":id", handler.GetByID)
		saleRoutes.DELETE(":id", handler.Delete)
	}
}
