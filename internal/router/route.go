package router

import (
	"ilanver/internal/api/category"
	"ilanver/internal/api/user"
	config "ilanver/internal/configs"

	_mware "ilanver/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteInit(e *echo.Echo) {

	e.POST("/login", user.Login)
	e.POST("/register", user.Register)

	admin := e.Group("")
	admin.Use(middleware.JWTWithConfig(config.JWTConfig))
	admin.Use(_mware.Auth)

	admincategory := admin.Group("")
	e.GET("/category/main", category.MainCategory)
	admincategory.POST("/category", category.Insert)
	admincategory.PUT("/category", category.Update)
	admincategory.DELETE("/category", category.Delete)

	adminUser := admin.Group("")
	adminUser.PUT("/user/update", user.Update)
	e.Start(":8080")
}
