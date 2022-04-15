package service

import (
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"ilanver/pkg/logger"
	"ilanver/request"
	"strconv"
)

type IProductStateService interface {
	GetByID(id string) (model.ProductState, error)
	GetAll() ([]model.ProductState, error)
	Insert(req request.InsertProductState) error
	Update(req request.UpdateProductState) error
	Delete(id string) error
}

type ProductStateRepository struct {
	repository repository.IProductStateRepository
}

func NewProductStateService(repository repository.IProductStateRepository) IProductStateService {
	return &ProductStateRepository{
		repository: repository,
	}
}

func (p *ProductStateRepository) GetByID(id string) (model.ProductState, error) {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warnf(4, "ProductStateRepository.GetByID: %v", err)
		return model.ProductState{}, err
	}

	productState, err := p.repository.GetByID(idInt)
	if err != nil {
		logger.Errorf(4, "ProductStateRepository.GetByID: %v", err)
	}
	return productState, err
}

func (p *ProductStateRepository) GetAll() ([]model.ProductState, error) {

	productStates, err := p.repository.GetAll()
	if err != nil {
		logger.Errorf(4, "ProductStateRepository.GetAll: %v", err)
	}
	return productStates, err
}

func (p *ProductStateRepository) Insert(req request.InsertProductState) error {
	productState := model.ProductState{
		State: req.State,
	}

	err := p.repository.Insert(&productState)
	if err != nil {
		logger.Errorf(4, "ProductStateRepository.Insert: %v", err)
	}
	return err
}

func (p *ProductStateRepository) Update(req request.UpdateProductState) error {
	productState, err := p.repository.GetByID(int(req.ID))
	if err != nil {
		logger.Errorf(4, "ProductStateRepository.Update: %v", err)
	}

	productState.State = req.State

	err = p.repository.Update(&productState)
	if err != nil {
		logger.Errorf(4, "ProductStateRepository.Update: %v", err)
	}

	return err
}

func (p *ProductStateRepository) Delete(id string) error {

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.Warnf(4, "ProductStateRepository.GetByID: %v", err)
		return err
	}

	err = p.repository.Delete(uint(idInt))
	if err != nil {
		logger.Errorf(4, "ProductStateRepository.Delete: %v", err)
	}

	return err
}
