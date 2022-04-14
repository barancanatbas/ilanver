package product

import (
	"ilanver/internal/handler"
	"ilanver/internal/repository"
	"ilanver/internal/service"

	"gorm.io/gorm"
)

func Init(tx *gorm.DB) handler.IProductHandler {
	productRepository := repository.NewProductRepository(tx)
	repository := repository.NewRepository(tx)
	productService := service.NewProductService(productRepository, repository)
	handler := handler.NewProductHandler(productService)

	return handler
}
