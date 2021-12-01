package user

import (
	"ilanver/internal/helpers"
	"ilanver/internal/models"
	"ilanver/repository"
	"ilanver/request"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {

	return c.String(200, "ok")
}

func Register(c echo.Context) error {
	var req request.UserRegister

	if helpers.Validator(&c, &req) != "" {
		return nil
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 4)

	existsPhone := repository.Get().User().ExistsPhone(req.Phone)
	if existsPhone {
		return c.String(200, "bu telefon zaten var")
	}
	user := models.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: string(password),
	}

	err := repository.Get().User().Register(user)
	// bu alanda user detail alanı da eklenecek
	if err != nil {
		return c.String(200, "Kayıt yapılamadı")
	}

	return c.String(200, "kayıt yapıldı.")
}
