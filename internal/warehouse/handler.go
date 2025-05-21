package warehouse

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

func NewHandler(uc UseCase) *Handler {
	return &Handler{uc}
}

func (h *Handler) RegisterRoutes(rg *gin.RouterGroup) {
	w := rg.Group("/warehouses")
	{
		w.POST("", h.Create)
		w.GET("", h.GetAll)
		w.GET("/:id", h.GetByID)
		w.PUT("/:id", h.Update)
		w.DELETE("/:id", h.Delete)
	}
}

func (h *Handler) Create(c *gin.Context) {
	var w Warehouse
	if err := c.ShouldBindJSON(&w); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.uc.Create(&w); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, w, "warehouse created")
}

func (h *Handler) GetAll(c *gin.Context) {
	list, err := h.uc.GetAll()
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to list warehouses", err)
		return
	}
	response.Success(c, list, "warehouses retrieved")
}

func (h *Handler) GetByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	w, err := h.uc.GetByID(uint(id64))
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch warehouse", err)
		return
	}
	if w == nil {
		response.Error(c, http.StatusNotFound, "warehouse not found")
		return
	}
	response.Success(c, w, "warehouse retrieved")
}

func (h *Handler) Update(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	var w Warehouse
	if err := c.ShouldBindJSON(&w); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	w.ID = uint(id64)
	updated, err := h.uc.Update(&w)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, updated, "warehouse updated")
}

func (h *Handler) Delete(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.uc.Delete(uint(id64)); err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to delete warehouse", err)
		return
	}
	c.Status(http.StatusNoContent)
}
