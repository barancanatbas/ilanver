package service

import (
	"errors"
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
	RepoUser       repository.UserRepo
	RepoAddress    repository.AddressRepo
	RepoUserDetail repository.IUserDetailRepo
}

var _ IUserService = UserService{}

func NewUserService(repoUser repository.UserRepo, repoAddress repository.AddressRepo, repoDetail repository.IUserDetailRepo) UserService {
	return UserService{
		RepoUser:       repoUser,
		RepoAddress:    repoAddress,
		RepoUserDetail: repoDetail,
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

	err := s.RepoAddress.Save(&address)
	if err != nil {
		return model.User{}, err
	}

	userDetail := model.UserDetail{
		Adressfk: address.ID,
	}

	err = s.RepoUserDetail.Save(&userDetail)
	if err != nil {
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

	err = s.RepoUser.Save(&user)

	return user, err
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

	err = s.RepoUser.Save(&user)
	return err
}
