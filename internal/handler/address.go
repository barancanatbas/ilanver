package handler

import (
	service "ilanver/internal/service"
	"ilanver/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAddressHandler interface {
	Update(c *gin.Context)
}

type AddressHandler struct {
	Service service.IAddressService
}

func NewAddressHandler(service service.IAddressService) IAddressHandler {
	return AddressHandler{
		Service: service,
	}
}

func (h AddressHandler) Update(c *gin.Context) {
	var req request.UpdateAddress

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.Update(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "başarılı"})
	return
}
