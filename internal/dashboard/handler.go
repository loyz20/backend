// internal/dashboard/handler.go
package dashboard

import (
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase UseCase
}

func NewHandler(usecase UseCase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) GetDashboard(c *gin.Context) {
	data, err := h.usecase.GetDashboardData()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, data, "dashboard data retrieved successfully")
}
