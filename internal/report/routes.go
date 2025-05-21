// internal/report/routes.go
package report

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, h *Handler) {
	rg.GET("/reports/purchase", h.PurchaseReport)
	rg.GET("/reports/sale", h.SaleReport)
	rg.GET("/reports/stock", h.StockReport)
	rg.GET("/reports/financial", h.FinancialReport)
}
