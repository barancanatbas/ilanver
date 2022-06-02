package main

import (
	config "ilanver/internal/config"
	"ilanver/internal/middleware"
	"ilanver/internal/queue"
	"ilanver/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	config.Migrate(config.DB)
	r := gin.Default()
	r.Use(middleware.Limitter())

	// şuanda microservis olmadığında veya başka bir servis yapımız olmadığı için bu şekilde kullandık.
	// bir product insert işlemi yapan bir service olsaydı bunu yazmayacaktık
	//go queue.ConsumeInsertProduct("insertProduct")
	go queue.ConsumeUpdateProduct("updateProduct")

	router.Init(config.DB, config.ElasticDB, r)

	r.Run()
}
