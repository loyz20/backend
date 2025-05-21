// internal/report/handler.go
package report

import (
	"backend/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase UseCase
}

func NewHandler(u UseCase) *Handler {
	return &Handler{u}
}

func (h *Handler) SaleReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	customerName := c.Query("customer")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid start_date format, use YYYY-MM-DD")
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid end_date format, use YYYY-MM-DD")
			return
		}
	}

	report, err := h.usecase.GenerateSaleReport(startDate, endDate, customerName)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to generate report")
		return
	}

	response.Success(c, report, "report generated successfully")
}

func (h *Handler) StockReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	sku := c.Query("sku")
	name := c.Query("name")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid date format, use YYYY-MM-DD")
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid end_date format, use YYYY-MM-DD")
			return
		}
	}

	data, err := h.usecase.GenerateStockReport(startDate, endDate, sku, name)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, data, "report generated successfully")
}

// Endpoint: GET /reports/purchase?start_date=2023-01-01&end_date=2023-12-31&supplier=abc
func (h *Handler) PurchaseReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	supplierName := c.Query("supplier")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-01-02", startDateStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid start_date format, use YYYY-MM-DD")
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-01-02", endDateStr)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid end_date format, use YYYY-MM-DD")
			return
		}
	}

	report, err := h.usecase.GeneratePurchaseReport(startDate, endDate, supplierName)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to generate report")
		return
	}

	response.Success(c, report, "report generated successfully")
}

func (h *Handler) FinancialReport(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, use YYYY-MM-DD"})
		return
	}
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, use YYYY-MM-DD"})
		return
	}

	report, err := h.usecase.GenerateFinancialReport(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate financial report"})
		return
	}

	c.JSON(http.StatusOK, report)
}
