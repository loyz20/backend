package user

import (
	"backend/internal/shared"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	uc := NewUseCase(repo)
	handler := NewHandler(uc)
	rg.POST("/register", handler.Register)
	rg.POST("/login", handler.Login)
	users := rg.Group("/auth/users")
	{
		users.Use(shared.JWTAuthMiddleware())
		users.GET("/", handler.GetAll)
		users.GET("/:id", handler.GetByID)
		users.PUT("/:id", handler.Update)
		users.DELETE("/:id", handler.Delete)
	}
}
