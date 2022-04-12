package handler

import (
	"ilanver/internal/helpers"
	service "ilanver/internal/service"
	"ilanver/request"

	"github.com/gin-gonic/gin"
)

type ICategoryHandler interface {
	GetAll(c *gin.Context)
	GetSubCategories(c *gin.Context)
	Insert(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CategoryHandler struct {
	service service.ICategoryService
}

func NewCategoryHandler(service service.ICategoryService) ICategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) GetAll(c *gin.Context) {
	page := c.Query("page")

	categories, err := h.service.GetAll(page)
	if err != nil {
		c.JSON(500, helpers.BasicError(400, err))
		return
	}
	c.JSON(200, helpers.BasicReturn(200, categories))
}

func (h *CategoryHandler) GetSubCategories(c *gin.Context) {
	id := c.Param("id")

	categories, err := h.service.GetSubCategories(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, helpers.BasicReturn(200, categories))
}

func (h *CategoryHandler) Insert(c *gin.Context) {
	var category request.InsertCategory
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = h.service.Insert(category)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, helpers.BasicReturn(200, "ekleme işlemi başarılı"))
}

func (h *CategoryHandler) Update(c *gin.Context) {
	var category request.UpdateCategory
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(500, err)
		return
	}
	err = h.service.Update(category)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, helpers.BasicReturn(200, "güncelleme işlemi başarılı"))
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, helpers.BasicReturn(200, "silme işlemi başarılı"))
}
