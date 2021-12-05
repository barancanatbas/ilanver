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
	e.GET("/category/main", category.MainCategory)

	admin := e.Group("")
	admin.Use(middleware.JWTWithConfig(config.JWTConfig))
	admin.Use(_mware.Auth)

	admincategory := admin.Group("")

	admincategory.POST("/category", category.Insert)
	e.Start(":8080")
}
