package main

import (
	configs "ilanver/internal/configs"
	"ilanver/internal/router"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	configs.Init()
	router.RouteInit(e)
}
