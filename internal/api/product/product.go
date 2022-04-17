package product

import (
	"ilanver/internal/handler"
	"ilanver/internal/repository"
	"ilanver/internal/service"

	"gorm.io/gorm"
)

func Init(tx *gorm.DB) handler.IProductHandler {
	repoProduct := repository.NewProductRepository(tx)
	repo := repository.NewRepository(tx)
	repoAddress := repository.NewAddressRepository(tx)
	productService := service.NewProductService(repoProduct, repo, repoAddress)
	handler := handler.NewProductHandler(productService)

	return handler
}
