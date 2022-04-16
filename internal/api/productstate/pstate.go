package productstate

import (
	"ilanver/internal/handler"
	"ilanver/internal/repository"
	"ilanver/internal/service"

	"gorm.io/gorm"
)

func Init(tx *gorm.DB) handler.IProductStateHandler {
	productStateRepository := repository.NewProductStateRepository(tx)
	productStateService := service.NewProductStateService(productStateRepository)
	handler := handler.NewProductStateHandler(productStateService)

	return handler
}
