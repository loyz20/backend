// internal/product/handler.go
package product

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

func NewHandler(u UseCase) *Handler {
	return &Handler{uc: u}
}

func (h *Handler) RegisterRoutes(r *gin.RouterGroup) {
	grp := r.Group("/products")
	grp.GET("", h.GetAll)
	grp.GET(":id", h.GetByID)
	grp.POST("", h.Create)
	grp.PUT(":id", h.Update)
	grp.DELETE(":id", h.Delete)
}

func (h *Handler) GetAll(c *gin.Context) {
	products, err := h.uc.GetAll()
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch products", err)
		return
	}
	response.Success(c, products, "products retrieved successfully")
}

func (h *Handler) GetByID(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}
	product, err := h.uc.GetByID(uint(id64))
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch product", err)
		return
	}
	if product == nil {
		response.Error(c, http.StatusNotFound, "product not found")
		return
	}
	response.Success(c, product, "product retrieved successfully")
}

func (h *Handler) Create(c *gin.Context) {
	var input Product
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.uc.Create(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, input, "product created successfully")
}

func (h *Handler) Update(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}

	var input Product
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	updated, err := h.uc.Update(uint(id64), &input)
	if err != nil {
		if err.Error() == "product not found" {
			response.Error(c, http.StatusNotFound, err.Error())
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	response.Success(c, updated, "product updated successfully")
}

func (h *Handler) Delete(c *gin.Context) {
	id64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}
	if err := h.uc.Delete(uint(id64)); err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to delete product", err)
		return
	}
	c.Status(http.StatusNoContent)
}
