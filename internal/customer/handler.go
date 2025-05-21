package customer

import (
	"backend/internal/shared"
	"backend/pkg/response"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase UseCase
}

func NewHandler(usecase UseCase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) GetAll(c *gin.Context) {
	customers, err := h.usecase.GetAll()
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch customers", err)
		return
	}
	response.Success(c, customers, "customers retrieved successfully")
}

func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}
	cust, err := h.usecase.GetByID(uint(id))
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch customer", err)
		return
	}
	if cust == nil {
		response.Error(c, http.StatusNotFound, "customer not found")
		return
	}
	response.Success(c, cust, "customer retrieved successfully")
}

func (h *Handler) Create(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
		Address     string `json:"address"`
		SIPA        string `json:"sipa"`
		ExpiredDate string `json:"expired_date"` // RFC3339 format (ex: 2025-12-31T00:00:00Z)
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var expired *time.Time
	if input.ExpiredDate != "" {
		t, err := time.Parse(time.RFC3339, input.ExpiredDate)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid expired_date format")
			return
		}
		expired = &t
	}

	customer := &Customer{
		Name:        input.Name,
		Phone:       input.Phone,
		Email:       input.Email,
		Address:     input.Address,
		SIPA:        input.SIPA,
		ExpiredDate: expired,
	}

	if err := h.usecase.Create(customer); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, customer, "customer created successfully")
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}

	var input struct {
		Name        string `json:"name" binding:"required"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
		Address     string `json:"address"`
		SIPA        string `json:"sipa"`
		ExpiredDate string `json:"expired_date"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var expired *time.Time
	if input.ExpiredDate != "" {
		t, err := time.Parse(time.RFC3339, input.ExpiredDate)
		if err != nil {
			response.Error(c, http.StatusBadRequest, "invalid expired_date format")
			return
		}
		expired = &t
	}

	customer := &Customer{
		Name:        input.Name,
		Phone:       input.Phone,
		Email:       input.Email,
		Address:     input.Address,
		SIPA:        input.SIPA,
		ExpiredDate: expired,
	}

	if err := h.usecase.Update(uint(id), customer); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, customer, "customer updated successfully")
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}
	if err := h.usecase.Delete(uint(id)); err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to delete customer", err)
		return
	}
	c.Status(http.StatusNoContent)
}
