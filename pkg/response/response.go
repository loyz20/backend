package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`            // "success" / "error"
	Message string      `json:"message,omitempty"` // Custom message
	Data    interface{} `json:"data,omitempty"`    // Optional payload
}

// Success response
func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error response
func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Status:  "error",
		Message: message,
	})
}
