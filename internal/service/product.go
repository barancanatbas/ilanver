package service

import (
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"ilanver/pkg/logger"
	"strconv"
)

type IProductService interface {
	GetByID(id string) (model.Product, error)
}

type ProductService struct {
	repoProduct repository.IProductRepository
	repository  repository.IRepository
}

func NewProductService(repoProduct repository.IProductRepository, repository repository.IRepository) IProductService {
	return ProductService{
		repoProduct: repoProduct,
		repository:  repository,
	}
}

func (p ProductService) GetByID(id string) (model.Product, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Errorf(4, "ProductService.GetByID: %v", err)
		return model.Product{}, err
	}

	product, err := p.repoProduct.GetByID(idInt)
	return product, err
}
