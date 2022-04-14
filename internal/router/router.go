package router

import (
	"ilanver/internal/api/address"
	"ilanver/internal/api/category"
	"ilanver/internal/api/product"
	"ilanver/internal/api/user"
	"ilanver/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(tx *gorm.DB, router *gin.Engine) {
	userHandler := user.Init(tx)
	addressHandler := address.Init(tx)
	categoryHandler := category.Init(tx)
	productHandler := product.Init(tx)

	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.Register)
	router.POST("/lost/password/confrim", userHandler.LostPasswordConfrim)
	router.POST("/lost/password", userHandler.ChangePasswordForCode)

	auth := router.Use(middleware.JWTAuthMiddleware(false, "secret"))

	// user routes
	auth.PUT("/user", userHandler.Update)
	auth.PUT("/user/address", addressHandler.Update)
	auth.PUT("/user/password", userHandler.ChangePassword)

	// category routes
	auth.GET("/categories", categoryHandler.GetAll)
	auth.GET("/categories/:id/sub", categoryHandler.GetSubCategories)
	auth.POST("/categories", categoryHandler.Insert)
	auth.PUT("/categories/:id", categoryHandler.Update)
	auth.DELETE("/categories/:id", categoryHandler.Delete)

	auth.GET("/product/:id", productHandler.GetByID)

}
