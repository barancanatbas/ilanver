package category

import (
	"ilanver/internal/handler"
	"ilanver/internal/repository"
	"ilanver/internal/service"

	"gorm.io/gorm"
)

func Init(tx *gorm.DB) handler.ICategoryHandler {
	repo := repository.NewCategoryRepository(tx)
	repository := repository.NewRepository(tx)
	service := service.NewCategoryService(repo, repository)
	handler := handler.NewCategoryHandler(service)

	return handler
}
