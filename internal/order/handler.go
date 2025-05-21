package order

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
	pos, err := h.usecase.GetAll()
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch purchase orders", err)
		return
	}
	response.Success(c, pos, "purchase orders retrieved successfully")
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	po, err := h.usecase.GetByID(uint(id))
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch purchase order", err)
		return
	}
	if po == nil {
		response.Error(c, http.StatusNotFound, "purchase order not found")
		return
	}
	response.Success(c, po, "purchase order retrieved successfully")
}

func (h *Handler) Create(c *gin.Context) {
	var input Order
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Optional: validate date or set default date now
	if input.OrderDate.IsZero() {
		input.OrderDate = time.Now()
	}

	if err := h.usecase.Create(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, input, "purchase order created successfully")
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input Order
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	input.ID = uint(id)

	if err := h.usecase.Update(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, input, "purchase order updated successfully")
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to delete purchase order", err)
		return
	}

	c.Status(http.StatusNoContent)
}
