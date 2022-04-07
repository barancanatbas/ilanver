package service

import (
	"errors"
	"fmt"
	"ilanver/internal/cache"
	"ilanver/internal/config"
	"ilanver/internal/helpers"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"ilanver/request"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(req request.UserLogin) (model.User, string, error)
	Register(req request.UserRegister) (model.User, error)
	Update(req request.UserUpdate) error
	LostPassword(c *gin.Context, req request.UserLostPassword) error
	ChangePasswordForCode(c *gin.Context, req request.UserChangePasswordForCode) error
	ChangePassword(req request.UserChangePassword) error
}

type UserService struct {
	Repository     repository.IRepository
	RepoUser       repository.IUserRepo
	RepoAddress    repository.IAddressRepo
	RepoUserDetail repository.IUserDetailRepo
}

func NewUserService(repoUser repository.IUserRepo, repoAddress repository.IAddressRepo, repoDetail repository.IUserDetailRepo, repository repository.IRepository) IUserService {
	return UserService{
		RepoUser:       repoUser,
		RepoAddress:    repoAddress,
		RepoUserDetail: repoDetail,
		Repository:     repository,
	}
}

func (s UserService) Login(req request.UserLogin) (model.User, string, error) {

	user, err := s.RepoUser.Login(req.Phone)
	if err != nil {
		return model.User{}, "", err
	}

	passwordControl := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if passwordControl != nil {
		return model.User{}, "", errors.New("şifre doğrulanmadı")
	}

	// özel oluşturulmuş bir struct tan bir nesne oluşturduk
	claims := &config.JwtCustom{
		User:          user,
		Authorization: 1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := Token.SignedString([]byte("secret"))
	if err != nil {
		return model.User{}, "", err
	}

	return user, t, nil
}

func (s UserService) Register(req request.UserRegister) (model.User, error) {

	address := model.Adress{
		Districtfk: req.Districtfk,
		Detail:     req.Description,
	}

	// create transaction for mysql
	tx := s.Repository.CreateTX()

	err := s.RepoAddress.WitchTX(tx).Save(&address)
	if err != nil {
		fmt.Println("girdi 1")
		s.Repository.RollBack()
		return model.User{}, err
	}

	userDetail := model.UserDetail{
		Adressfk: address.ID,
	}

	err = s.RepoUserDetail.WitchTX(tx).Save(&userDetail)
	if err != nil {
		fmt.Println("girdi 2")
		s.Repository.RollBack()
		return model.User{}, err
	}
	date, _ := time.Parse("02.01.2006", req.Birthday)

	user := model.User{
		Name:         req.Name,
		Surname:      req.Surname,
		Phone:        req.Phone,
		Email:        req.Email,
		UserDetailfk: userDetail.ID,
		Birthday:     date,
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 4)

	user.Password = string(password)

	err = s.RepoUser.WitchTX(tx).Save(&user)
	if err != nil {
		fmt.Println("girdi 3")
		s.Repository.RollBack()
		return model.User{}, err
	}
	s.Repository.Commit()

	return user, nil
}

func (s UserService) Update(req request.UserUpdate) error {

	user, err := s.RepoUser.Get(req.ID)
	if err != nil {
		return err
	}

	user.Name = req.Name
	user.Phone = req.Phone
	user.Surname = req.Surname
	user.Email = req.Email
	bDate, _ := time.Parse("02.01.2006", req.Birthday)
	user.Birthday = bDate

	err = s.RepoUser.Update(&user)
	return err
}

func (s UserService) LostPassword(c *gin.Context, req request.UserLostPassword) error {

	code := helpers.RandNumber(1000, 9999)
	ip := c.ClientIP()

	key := "lostPassword:" + req.Phone

	// burada bir mail veya sms gönderildiğini düşünülebilir

	cache.SetHashCache(key, map[string]string{
		"code":    strconv.Itoa(code),
		"phone":   req.Phone,
		"ip":      ip,
		"confirm": "false",
	})

	return nil
}

func (s UserService) ChangePasswordForCode(c *gin.Context, req request.UserChangePasswordForCode) error {

	key := "lostPassword:" + req.Phone
	ip := c.ClientIP()

	if !cache.Exists(key) {
		return errors.New("Lütfen kayıtlı bir kodunuzu giriniz")
	}

	data := cache.GetHashCache(key)

	if data["code"] != req.Code {
		return errors.New("kod doğrulanmadı")
	}

	if data["ip"] != ip {
		return errors.New("kod doğrulanmadı")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 4)

	user, err := s.RepoUser.GetByPhone(req.Phone)
	if err != nil {
		return err
	}

	user.Password = string(password)

	err = s.RepoUser.Update(&user)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) ChangePassword(req request.UserChangePassword) error {

	auth := helpers.AuthUser

	user, err := s.RepoUser.Get(auth.ID)
	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 4)

	user.Password = string(password)

	err = s.RepoUser.Update(&user)
	if err != nil {
		return err
	}

	return nil
}
