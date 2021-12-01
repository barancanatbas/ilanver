package config

import (
	"ilanver/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustom struct {
	User models.User `json:"user"`
	jwt.StandardClaims
}

var JWTConfig = middleware.JWTConfig{
	Claims:     &JwtCustom{},
	SigningKey: []byte("secret"),
}
