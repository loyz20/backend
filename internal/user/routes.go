package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	uc := NewUseCase(repo)
	handler := NewHandler(uc)

	users := rg.Group("/users")
	{
		users.POST("/register", handler.Register)
		users.POST("/login", handler.Login)
		users.GET("/", handler.GetAll)
		users.GET("/:id", handler.GetByID)
		users.PUT("/:id", handler.Update)
		users.DELETE("/:id", handler.Delete)
	}
}
