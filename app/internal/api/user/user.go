package user

import (
	handler "ilanver/internal/handler"
	"ilanver/internal/repository"
	service "ilanver/internal/service"

	"gorm.io/gorm"
)

func Init(tx *gorm.DB) handler.IUserHandler {
	repoUser := repository.NewUserRepository(tx)
	repoAddress := repository.NewAddressRepository(tx)
	repoDetail := repository.NewUserDetailRepository(tx)
	repository := repository.NewRepository(tx)

	service := service.NewUserService(repoUser, repoAddress, repoDetail, repository)
	handler := handler.NewUserHandler(service)

	return handler
}
