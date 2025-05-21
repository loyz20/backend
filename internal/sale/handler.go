package sale

import (
	"net/http"
	"strconv"
	"time"

	"backend/internal/shared"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase UseCase
}

func NewHandler(usecase UseCase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) GetAll(c *gin.Context) {
	sos, err := h.usecase.GetAll()
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch sale orders", err)
		return
	}
	response.Success(c, sos, "sale orders retrieved successfully")
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	so, err := h.usecase.GetByID(uint(id))
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch sale order", err)
		return
	}
	if so == nil {
		response.Error(c, http.StatusNotFound, "sale order not found")
		return
	}
	response.Success(c, so, "sale order retrieved successfully")
}

func (h *Handler) Create(c *gin.Context) {
	var input Sale
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if input.OrderDate.IsZero() {
		input.OrderDate = time.Now()
	}

	if err := h.usecase.Create(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, input, "sale order created successfully")
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input Sale
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	input.ID = uint(id)

	if err := h.usecase.Update(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, input, "sale order updated successfully")
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to delete sale order", err)
		return
	}

	c.Status(http.StatusNoContent)
}
