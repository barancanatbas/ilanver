package main

import (
	configs "ilanver/internal/configs"
	"ilanver/internal/router"
	"ilanver/repository"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "ilanver/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.sw
func main() {

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	configs.Init()
	repository.Set()
	router.RouteInit(e)
}
