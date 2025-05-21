package shared

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, status int, message string, err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
	c.JSON(status, gin.H{
		"status":  "error",
		"message": message,
	})
}
