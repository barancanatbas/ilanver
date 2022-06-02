package handler

import (
	"ilanver/internal/helpers"
	service "ilanver/internal/service"
	"ilanver/request"

	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	//GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Save(c *gin.Context)
	Update(c *gin.Context)
}

type ProductHandler struct {
	service service.IProductService
}

func NewProductHandler(service service.IProductService) IProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	product, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, helpers.BasicReturn(200, product))
}

func (h *ProductHandler) Save(c *gin.Context) {
	var product request.InsertProduct
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = h.service.Save(product)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, helpers.BasicReturn(200, "Başarılı."))
}

func (h *ProductHandler) Update(c *gin.Context) {
	var product request.UpdateProduct
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(500, err)
		return
	}

	err = h.service.Update(product)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, helpers.BasicReturn(200, "Başarılı."))
}
