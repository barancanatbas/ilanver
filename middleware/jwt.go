package middleware

import (
	"ilanver/repository"

	config "ilanver/internal/configs"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*config.JwtCustom)

		// vt kontrolü
		_, err := repository.Get().User().Login(claims.User)
		if err != nil {
			return c.JSON(200, "Bilinmeyen bir hata oluştu")
		}

		return next(c)
	}
}
