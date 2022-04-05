package service

import (
	"errors"
	"fmt"
	"ilanver/internal/config"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"ilanver/request"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(req request.UserLogin) (model.User, string, error)
	Register(req request.UserRegister) (model.User, error)
	Update(req request.UserUpdate) error
}

type UserService struct {
	Repository     repository.IRepository
	RepoUser       repository.IUserRepo
	RepoAddress    repository.IAddressRepo
	RepoUserDetail repository.IUserDetailRepo
}

var _ IUserService = UserService{}

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

	err := s.RepoAddress.Save(&address, tx)
	if err != nil {
		fmt.Println("girdi 1")
		s.Repository.RollBack()
		return model.User{}, err
	}

	userDetail := model.UserDetail{
		Adressfk: address.ID,
	}

	err = s.RepoUserDetail.Save(&userDetail, tx)
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

	err = s.RepoUser.Save(&user, tx)
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
