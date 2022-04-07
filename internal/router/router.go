package router

import (
	"ilanver/internal/api/address"
	"ilanver/internal/api/user"
	"ilanver/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(tx *gorm.DB, router *gin.Engine) {
	userHandler := user.Init(tx)
	addressHandler := address.Init(tx)

	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.Register)
	router.POST("/lost/password/confrim", userHandler.LostPasswordConfrim)
	router.POST("/lost/password", userHandler.ChangePasswordForCode)

	auth := router.Use(middleware.JWTAuthMiddleware(false, "secret"))

	auth.PUT("/user", userHandler.Update)
	auth.PUT("/user/address", addressHandler.Update)
	auth.PUT("/user/password", userHandler.ChangePassword)

}
