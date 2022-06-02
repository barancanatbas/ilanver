package address

import (
	"ilanver/internal/handler"
	"ilanver/internal/repository"
	"ilanver/internal/service"

	"gorm.io/gorm"
)

func Init(tx *gorm.DB) handler.IAddressHandler {
	repoAddress := repository.NewAddressRepository(tx)
	repository := repository.NewRepository(tx)

	service := service.NewAddressService(repoAddress, repository)
	handler := handler.NewAddressHandler(service)

	return handler
}
