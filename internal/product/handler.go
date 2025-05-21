package product

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

// GetAll returns all products with their batches
func (h *Handler) GetAll(c *gin.Context) {
	products, err := h.usecase.GetAll()
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch products", err)
		return
	}
	response.Success(c, products, "products retrieved successfully")
}

// GetByID returns single product by id with batches
func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}
	prod, err := h.usecase.GetByID(uint(id))
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch product", err)
		return
	}
	if prod == nil {
		response.Error(c, http.StatusNotFound, "product not found")
		return
	}
	response.Success(c, prod, "product retrieved successfully")
}

// Create a new product (no batch here, batch handled separately)
func (h *Handler) Create(c *gin.Context) {
	var input struct {
		Name          string  `json:"name" binding:"required"`
		SKU           string  `json:"sku" binding:"required"`
		Category      string  `json:"category"`
		Unit          string  `json:"unit"`
		PurchasePrice float64 `json:"purchase_price"`
		SellingPrice  float64 `json:"selling_price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	prod := &Product{
		Name:          input.Name,
		SKU:           input.SKU,
		Category:      input.Category,
		Unit:          input.Unit,
		PurchasePrice: input.PurchasePrice,
		SellingPrice:  input.SellingPrice,
	}

	if err := h.usecase.Create(prod); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, prod, "product created successfully")
}

// Update product info (excluding batches)
func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}

	var input struct {
		Name          string  `json:"name" binding:"required"`
		SKU           string  `json:"sku" binding:"required"`
		Category      string  `json:"category"`
		Unit          string  `json:"unit"`
		PurchasePrice float64 `json:"purchase_price"`
		SellingPrice  float64 `json:"selling_price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	updateProd := &Product{
		Name:          input.Name,
		SKU:           input.SKU,
		Category:      input.Category,
		Unit:          input.Unit,
		PurchasePrice: input.PurchasePrice,
		SellingPrice:  input.SellingPrice,
	}

	updated, err := h.usecase.Update(uint(id), updateProd)
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

// Delete product by id
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid id format")
		return
	}

	if err := h.usecase.Delete(uint(id)); err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to delete product", err)
		return
	}

	c.Status(http.StatusNoContent)
}

// -- Batch Handlers --

// GetBatchesByProductID returns all batches for a product
func (h *Handler) GetBatchesByProductID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid product id format")
		return
	}
	batches, err := h.usecase.GetBatchesByProductID(uint(id))
	if err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to fetch batches", err)
		return
	}
	response.Success(c, batches, "batches retrieved successfully")
}

// CreateBatch for a product
func (h *Handler) CreateBatch(c *gin.Context) {
	var input struct {
		ProductID   uint   `json:"product_id" binding:"required"`
		BatchNumber string `json:"batch_number" binding:"required"`
		Quantity    int    `json:"quantity" binding:"required"`
		ExpiryDate  string `json:"expiry_date" binding:"required"` // expect RFC3339 string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	expiry, err := time.Parse(time.RFC3339, input.ExpiryDate)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid expiry_date format")
		return
	}

	batch := &ProductBatch{
		ProductID:         input.ProductID,
		BatchNumber:       input.BatchNumber,
		Quantity:          input.Quantity,
		RemainingQuantity: input.Quantity,
		ExpiryDate:        expiry,
	}

	if err := h.usecase.CreateBatch(batch); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, batch, "batch created successfully")
}

// UpdateBatch info
func (h *Handler) UpdateBatch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("batch_id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid batch id format")
		return
	}

	var input struct {
		BatchNumber       string `json:"batch_number" binding:"required"`
		Quantity          int    `json:"quantity" binding:"required"`
		RemainingQuantity int    `json:"remaining_quantity" binding:"required"`
		ExpiryDate        string `json:"expiry_date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	expiry, err := time.Parse(time.RFC3339, input.ExpiryDate)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid expiry_date format")
		return
	}

	batch := &ProductBatch{
		ID:                uint(id),
		BatchNumber:       input.BatchNumber,
		Quantity:          input.Quantity,
		RemainingQuantity: input.RemainingQuantity,
		ExpiryDate:        expiry,
	}

	if err := h.usecase.UpdateBatch(batch); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, batch, "batch updated successfully")
}

// DeleteBatch by id
func (h *Handler) DeleteBatch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("batch_id"), 10, 32)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid batch id format")
		return
	}

	if err := h.usecase.DeleteBatch(uint(id)); err != nil {
		shared.HandleError(c, http.StatusInternalServerError, "failed to delete batch", err)
		return
	}

	c.Status(http.StatusNoContent)
}
