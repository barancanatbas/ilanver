package service

import (
	"encoding/json"
	"ilanver/internal/helpers"
	"ilanver/internal/model"
	"ilanver/internal/queue"
	"ilanver/internal/repository"
	"ilanver/pkg/logger"
	"ilanver/request"
	"strconv"
)

type IProductService interface {
	GetByID(id string) (model.Product, error)
	Save(req request.InsertProduct) error
	Update(req request.UpdateProduct) error
}

type ProductService struct {
	repoProduct repository.IProductRepository
	repository  repository.IRepository
	repoAddress repository.IAddressRepo
}

func NewProductService(repoProduct repository.IProductRepository, repository repository.IRepository, repoAddress repository.IAddressRepo) IProductService {
	return ProductService{
		repoProduct: repoProduct,
		repository:  repository,
		repoAddress: repoAddress,
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

func (p ProductService) Save(req request.InsertProduct) error {

	tx := p.repository.CreateTX()
	user := helpers.AuthUser

	address := model.Adress{
		Detail:     req.AddressDetail,
		Districtfk: uint(req.Districtfk),
	}

	err := p.repoAddress.WithTx(tx).Save(&address)
	if err != nil {
		p.repository.RollBack()
		logger.Errorf(4, "ProductService.Save: %v", err)
		return err
	}

	pDetail := model.ProductDetail{
		Adressfk:       address.ID,
		ProductStatefk: uint(req.ProductStateFk),
		Description:    req.ProductDescription,
		Price:          req.Price,
		Categoryfk:     uint(req.CategoryFk),
	}

	err = p.repoProduct.WithTx(tx).SaveDetail(&pDetail)
	if err != nil {
		p.repository.RollBack()
		logger.Errorf(4, "ProductService.Save: %v", err)
		return err
	}

	product := model.Product{
		Title:           req.Title,
		Userfk:          user.ID,
		ProductDetailfk: pDetail.ID,
	}

	err = p.repoProduct.WithTx(tx).Save(&product)
	if err != nil {
		p.repository.RollBack()
		logger.Errorf(4, "ProductService.Save: %v", err)
		return err
	}

	// burada producttan yani bir nesne oluştur ve elasticsearche kaydetmek için rabbitmq ya gönder.
	productElastic := model.ProductElastic{
		ID:          product.ID,
		Title:       product.Title,
		Description: pDetail.Description,
		Price:       pDetail.Price,
	}

	data, err := json.Marshal(productElastic)
	if err != nil {
		logger.Warnf(4, "ProductService.Save: %v", err)
	}

	err = queue.NewQueue().Publish("insertProduct", data)
	if err != nil {
		logger.Warnf(4, "ProductService.Save: %v", err)
	}

	p.repository.Commit()

	return nil
}

func (p ProductService) Update(req request.UpdateProduct) error {

	product, err := p.repoProduct.GetByID(int(req.ID))

	tx := p.repository.CreateTX()

	product.Title = req.Title

	err = p.repoProduct.WithTx(tx).Save(&product)
	if err != nil {
		p.repository.RollBack()
		logger.Errorf(4, "ProductService.Update: %v", err)
		return err
	}

	product.ProductDetail.Description = req.ProductDescription
	product.ProductDetail.Price = req.Price
	product.ProductDetail.ProductStatefk = uint(req.ProductStateFk)
	product.ProductDetail.Categoryfk = uint(req.CategoryFk)

	err = p.repoProduct.WithTx(tx).SaveDetail(&product.ProductDetail)
	if err != nil {
		p.repository.RollBack()
		logger.Errorf(4, "ProductService.Update: %v", err)
		return err
	}

	p.repository.Commit()

	// burada rabbitmq ya istek atacak ve elasticsearchdeki verileri güncelleyecek.

	productElastic := model.ProductElastic{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.ProductDetail.Description,
		Price:       product.ProductDetail.Price,
	}

	data, err := json.Marshal(productElastic)
	if err != nil {
		logger.Warnf(4, "ProductService.Save: %v", err)
	}

	err = queue.NewQueue().Publish("updateProduct", data)
	if err != nil {
		logger.Warnf(4, "ProductService.Save: %v", err)
	}

	return nil
}
