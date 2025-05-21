// internal/supplier/handler.go
package supplier

import (
	"net/http"
	"strconv"

	"backend/internal/shared"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc UseCase
}

// NewHandler returns a new Supplier HTTP handler
func NewHandler(uc UseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetAll(c *gin.Context) {
	suppliers, err := h.uc.GetAll()
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch suppliers", err)
		return
	}
	response.Success(c, suppliers, "suppliers retrieved successfully")
}

func (h *Handler) GetByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}
	s, err := h.uc.GetByID(uint(id64))
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch supplier", err)
		return
	}
	if s == nil {
		response.Error(c, http.StatusNotFound, "supplier not found")
		return
	}
	response.Success(c, s, "supplier retrieved successfully")
}

func (h *Handler) Create(c *gin.Context) {
	var input Supplier
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.uc.Create(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, input, "supplier created successfully")
}

func (h *Handler) Update(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}
	var input Supplier
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	updated, err := h.uc.Update(uint(id64), &input)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, updated, "supplier updated successfully")
}

func (h *Handler) Delete(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}
	if err := h.uc.Delete(uint(id64)); err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to delete supplier", err)
		return
	}
	c.Status(http.StatusNoContent)
}
