package router

import (
	"ilanver/internal/api/user"
	"ilanver/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(tx *gorm.DB, router *gin.Engine) {
	userHandler := user.Init(tx)

	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.Register)

	auth := router.Use(middleware.JWTAuthMiddleware(false, "secret"))

	auth.PUT("/user", userHandler.Update)

}
