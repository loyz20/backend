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

	users := rg.Group("/auth")
	{
		users.POST("/register", handler.Register)
		users.POST("/login", handler.Login)
		auth := rg.Group("/auth/users")
		{
			auth.Use(shared.JWTAuthMiddleware())
			auth.GET("/", handler.GetAll)
			auth.GET("/:id", handler.GetByID)
			auth.PUT("/:id", handler.Update)
			auth.DELETE("/:id", handler.Delete)
		}
	}
}
