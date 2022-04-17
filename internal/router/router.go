package router

import (
	"ilanver/internal/api/address"
	"ilanver/internal/api/category"
	"ilanver/internal/api/product"
	"ilanver/internal/api/productstate"
	"ilanver/internal/api/user"
	"ilanver/internal/config"
	"ilanver/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(tx *gorm.DB, elasticDb *config.ElasticSearch, router *gin.Engine) {
	userHandler := user.Init(tx)
	addressHandler := address.Init(tx)
	categoryHandler := category.Init(tx)
	productHandler := product.Init(tx)
	productState := productstate.Init(tx)

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

	// product state routes
	auth.POST("/product/state", productState.Insert)
	auth.PUT("/product/state/:id", productState.Update)
	auth.DELETE("/product/state/:id", productState.Delete)
	auth.GET("/product/states", productState.GetAll)
	auth.GET("/product/states/:id", productState.Get)

	auth.GET("/product/:id", productHandler.GetByID)
	auth.POST("/product", productHandler.Save)
	auth.PUT("/product", productHandler.Update)

}
