package config

import (
	"ilanver/internal/model"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustom struct {
	User          model.User `json:"user"`
	Authorization uint       `json:"authorization"` // 1 = admin, 2 = sınırlı kullanıcı
	jwt.StandardClaims
}

// var JWTConfig = middleware.JWTConfig{
// 	Claims:     &JwtCustom{},
// 	SigningKey: []byte("secret"),
// }
