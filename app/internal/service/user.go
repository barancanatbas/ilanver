package service

import (
	"errors"
	"ilanver/internal/cache"
	"ilanver/internal/config"
	"ilanver/internal/helpers"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"ilanver/pkg/logger"
	"ilanver/request"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(req request.UserLogin) (model.User, string, error)
	Register(req request.UserRegister) (model.User, error)
	Update(req request.UserUpdate) error
	LostPassword(ip string, req request.UserLostPassword) error
	ChangePasswordForCode(ip string, req request.UserChangePasswordForCode) error
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
		logger.Errorf(4, "UserService.Login: %s", err.Error())
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
		logger.Errorf(4, "UserService.Login: %s", err.Error())
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

	err := s.RepoAddress.WithTx(tx).Save(&address)
	if err != nil {
		logger.Errorf(4, "UserService.Register: %s", err.Error())
		s.Repository.RollBack()
		return model.User{}, err
	}

	userDetail := model.UserDetail{
		Adressfk: address.ID,
	}

	err = s.RepoUserDetail.WithTx(tx).Save(&userDetail)
	if err != nil {
		logger.Errorf(4, "UserService.Register: %s", err.Error())
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

	err = s.RepoUser.WithTx(tx).Save(&user)
	if err != nil {
		logger.Errorf(4, "UserService.Register: %s", err.Error())
		s.Repository.RollBack()
		return model.User{}, err
	}
	s.Repository.Commit()

	return user, nil
}

func (s UserService) Update(req request.UserUpdate) error {

	user, err := s.RepoUser.Get(req.ID)
	if err != nil {
		logger.Errorf(4, "UserService.Update: %s", err.Error())
		return err
	}

	user.Name = req.Name
	user.Phone = req.Phone
	user.Surname = req.Surname
	user.Email = req.Email
	bDate, _ := time.Parse("02.01.2006", req.Birthday)
	user.Birthday = bDate

	err = s.RepoUser.Update(&user)
	if err != nil {
		logger.Errorf(4, "UserService.Update: %s", err.Error())
	}
	return err
}

func (s UserService) LostPassword(ip string, req request.UserLostPassword) error {

	code := helpers.RandNumber(1000, 9999)

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

func (s UserService) ChangePasswordForCode(ip string, req request.UserChangePasswordForCode) error {

	key := "lostPassword:" + req.Phone

	if !cache.Exists(key) {
		logger.Warnf(4, "UserService.ChangePasswordForCode: %s", "kod geçersiz")
		return errors.New("Lütfen kayıtlı bir kodunuzu giriniz")
	}

	data := cache.GetHashCache(key)

	if data["code"] != req.Code {
		logger.Warnf(4, "UserService.ChangePasswordForCode: %s", "kod doğru değil")
		return errors.New("kod doğrulanmadı")
	}

	if data["ip"] != ip {
		logger.Warnf(4, "UserService.ChangePasswordForCode: %s", "ip doğru değil")
		return errors.New("kod doğrulanmadı")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 4)

	user, err := s.RepoUser.GetByPhone(req.Phone)
	if err != nil {
		logger.Errorf(4, "UserService.ChangePasswordForCode: %s", err.Error())
		return err
	}

	user.Password = string(password)

	err = s.RepoUser.Update(&user)
	if err != nil {
		logger.Errorf(4, "UserService.ChangePasswordForCode: %s", err.Error())
		return err
	}

	return nil
}

func (s UserService) ChangePassword(req request.UserChangePassword) error {

	auth := helpers.AuthUser

	user, err := s.RepoUser.Get(auth.ID)
	if err != nil {
		logger.Errorf(4, "UserService.ChangePassword: %s", err.Error())
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 4)

	if err != nil {
		logger.Errorf(4, "UserService.ChangePassword: %s", err.Error())
		return err
	}

	user.Password = string(password)

	err = s.RepoUser.Update(&user)
	if err != nil {
		logger.Errorf(4, "UserService.ChangePassword: %s", err.Error())
		return err
	}

	return nil
}
