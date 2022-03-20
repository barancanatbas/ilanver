package user

import (
	handler "ilanver/internal/handler/handler_user"
	"ilanver/internal/repository"
	service "ilanver/internal/service/service_user"

	"gorm.io/gorm"
)

func Init(tx *gorm.DB) handler.UserHandler {
	repoUser := repository.NewUserRepository(tx)
	repoAddress := repository.NewAddressRepository(tx)
	repoDetail := repository.NewUserDetailRepository(tx)

	service := service.NewUserService(repoUser, repoAddress, repoDetail)
	handler := handler.NewUserHandler(service)

	return handler
}
