package handler

import (
	"ilanver/internal/helpers"
	"ilanver/internal/service"
	"ilanver/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IProductStateHandler interface {
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Insert(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ProductStateHandler struct {
	service service.IProductStateService
}

func NewProductStateHandler(service service.IProductStateService) IProductStateHandler {
	return &ProductStateHandler{
		service: service,
	}
}

func (p *ProductStateHandler) Get(c *gin.Context) {
	id := c.Param("id")
	productState, err := p.service.GetByID(id)
	if err != nil {
		c.JSON(500, helpers.BasicError(http.StatusBadRequest, err))
		return
	}
	c.JSON(200, helpers.BasicReturn(http.StatusOK, productState))
}

func (p *ProductStateHandler) GetAll(c *gin.Context) {
	productStates, err := p.service.GetAll()
	if err != nil {
		c.JSON(500, helpers.BasicError(http.StatusBadRequest, err))
		return
	}
	c.JSON(200, helpers.BasicReturn(http.StatusOK, productStates))
}

func (p *ProductStateHandler) Insert(c *gin.Context) {
	var req request.InsertProductState

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, helpers.BasicError(http.StatusBadRequest, err))
		return
	}

	err := p.service.Insert(req)
	if err != nil {
		c.JSON(500, helpers.BasicError(http.StatusBadRequest, err))
		return
	}

	c.JSON(200, helpers.BasicReturn(http.StatusOK, nil))
}

func (p *ProductStateHandler) Update(c *gin.Context) {
	var req request.UpdateProductState

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(500, helpers.BasicError(http.StatusBadRequest, err))
		return
	}

	err := p.service.Update(req)
	if err != nil {
		c.JSON(500, helpers.BasicError(http.StatusBadRequest, err))
		return
	}

	c.JSON(200, helpers.BasicReturn(http.StatusOK, nil))
}

func (p *ProductStateHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := p.service.Delete(id)
	if err != nil {
		c.JSON(500, helpers.BasicError(http.StatusBadRequest, err))
		return
	}

	c.JSON(200, helpers.BasicReturn(http.StatusOK, nil))
}
