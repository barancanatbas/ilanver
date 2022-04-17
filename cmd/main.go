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
	go queue.ConsumeInsertProduct("insertProduct")

	router.Init(config.DB, config.ElasticDB, r)

	r.Run()
}

//datadog
// tracer.Start(
// 	tracer.WithService("ilanver-go"),
// 	tracer.WithEnv("prod"),
// )
// defer tracer.Stop()

// if err := profiler.Start(
// 	profiler.WithService("ilanver-go"),
// 	profiler.WithEnv("prod"),
// 	profiler.WithProfileTypes(
// 		profiler.CPUProfile,
// 		profiler.HeapProfile,

// 		// The profiles below are disabled by
// 		// default to keep overhead low, but
// 		// can be enabled as needed.
// 		// profiler.BlockProfile,
// 		// profiler.MutexProfile,
// 		// profiler.GoroutineProfile,
// 	),
// ); err != nil {
// 	log.Fatal(err)
// }
// defer profiler.Stop()
