package user

import (
	config "ilanver/internal/configs"
	"ilanver/internal/helpers"
	"ilanver/internal/models"
	"ilanver/repository"
	"ilanver/request"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Login
// @Description Üyelerin giriş yapmasını sağlar
// @Tags user
// @Param body body request.UserLogin false " "
// @Router /Login [post]
func Login(c echo.Context) error {
	var req request.UserLogin
	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	user := models.User{
		Phone:    req.Phone,
		Password: req.Password,
	}
	err := repository.Get().User().Login(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "kullanıcı bulunamadı"))
	}

	passwordControl := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if passwordControl != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Şifre doğrulanma"))
	}

	claims := &config.JwtCustom{
		User: *&user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Token oluşturulamadı"))
	}

	return c.JSON(http.StatusBadRequest, echo.Map{"token": t, "user": user})
}

// @Summary Register
// @Description Üyelerin kayıt yapmasını sağlar adres bilgisini kayıt eder, user detay bilgilerini kayıt eder.
// @Tags user
// @Param body body request.UserRegister false " "
// @Router /register [post]
func Register(c echo.Context) error {
	var req request.UserRegister

	if helpers.Validator(&c, &req) != nil {
		return nil
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 4)

	existsPhone := repository.Get().User().ExistsPhone(req.Phone, 0)
	if existsPhone {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "bu telefon numarası kullanılmaktadır."})
	}

	// address save
	address := models.Adress{
		Detail:     req.Description,
		Districtfk: req.Districtfk,
	}
	err := repository.Get().Address().Save(&address)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "adres kayıt edilmedi"})
	}

	// save user detail
	detail := models.UserDetail{
		ProfilePhoto: "",
		Adressfk:     address.ID,
		Birthday:     req.Birthday,
	}

	err = repository.Get().UserDetail().Save(&detail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "detaylar katıt edilemedi"))
	}
	// register user
	user := models.User{
		Name:         req.Name,
		Phone:        req.Phone,
		Email:        req.Email,
		Password:     string(password),
		UserDetailfk: detail.ID,
	}
	err = repository.Get().User().Register(user)
	// TODO: burada user detail ve adres işlemlerini yapacaz
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "kayıt yapılamadı"})
	}
	return c.JSON(http.StatusOK, helpers.Response(nil, "oluşturma başarılı"))
}

// @Summary Update
// @Description üyenin bilgilerini günceller
// @Tags user
// @Param body body request.UserUpdate false " "
// @Router /user/update [put]
func Update(c echo.Context) error {
	var req request.UserUpdate
	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	user := helpers.Auth(&c)

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Surname != "" {
		user.Surname = req.Surname
	}
	if req.Phone != "" {
		dublicatePhone := repository.Get().User().ExistsPhone(req.Phone, user.ID)
		if dublicatePhone {
			return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Bu telefon kullanılmakta."))
		}
		user.Phone = req.Phone
	}
	if req.Email != "" {
		dublicateEmail := repository.Get().User().ExistsEmail(req.Email, user.ID)
		if !dublicateEmail {
			return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Bu E-mail kullanılmakta."))
		}
		user.Email = req.Email
	}
	if req.Description != "" {
		user.UserDetail.Adress.Detail = req.Description
	}
	if req.Birthday != "" {
		user.UserDetail.Birthday = req.Birthday
	}
	if req.Districtfk != 0 {
		dublicateDist := repository.Get().Address().ExistsDistric(req.Districtfk)
		if !dublicateDist {
			return c.JSON(http.StatusBadRequest, helpers.Response(nil, "İlçe bilgisi bulunamadı."))
		}
		user.UserDetail.Adress.Districtfk = req.Districtfk
	}

	err := repository.Get().User().Update(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(err, "hata var"))
	}
	return c.JSON(200, user)
}
