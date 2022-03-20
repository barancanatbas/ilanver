package handler

import (
	service "ilanver/internal/service/service_user"
	"ilanver/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Update(c *gin.Context)
}

type UserHandler struct {
	Service service.UserService
}

var _ IUserHandler = UserHandler{}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{
		Service: service,
	}
}

func (h UserHandler) Login(c *gin.Context) {
	var req request.UserLogin

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, token, err := h.Service.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
	return
}

func (h UserHandler) Register(c *gin.Context) {
	var req request.UserRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "message": "başarılı"})
	return
}

func (h UserHandler) Update(c *gin.Context) {
	var req request.UserUpdate

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.Update(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "güncelleme başarılı"})
	return
}
