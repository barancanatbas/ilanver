package router

import (
	"ilanver/internal/api/user"

	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Echo) {

	e.POST("/login", user.Login)
	e.POST("/register", user.Register)
	e.Start(":8080")
}
