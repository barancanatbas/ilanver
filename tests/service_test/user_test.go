package servicetest

import (
	"ilanver/internal/config"
	"ilanver/internal/helpers"
	"ilanver/internal/repository"
	"ilanver/internal/service"
	"ilanver/request"
	"strconv"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUserLogin(t *testing.T) {
	config.InitTest()

	repoUser := repository.NewUserRepository(config.DBTest)
	repoAddress := repository.NewAddressRepository(config.DBTest)
	repoDetail := repository.NewUserDetailRepository(config.DBTest)
	repository := repository.NewRepository(config.DBTest)

	service := service.NewUserService(repoUser, repoAddress, repoDetail, repository)

	phone := helpers.RandNumber(1000000000, 9999999997)
	phoneString := strconv.Itoa(phone)
	t.Log(phoneString)

	user, err := service.Register(request.UserRegister{
		Phone:       phoneString,
		Password:    "12345678",
		Name:        "test",
		Surname:     "test",
		Email:       "deneme@gmail.com",
		Birthday:    "01.01.2000",
		Districtfk:  1,
		Description: "test",
	})

	assert.Equal(t, user.Phone, phoneString)
	assert.Equal(t, err, nil)
	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Surname, "test")
	assert.Equal(t, user.Email, "deneme@gmail.com")

	userLogin, _, err := service.Login(request.UserLogin{
		Phone:    phoneString,
		Password: "12345678",
	})

	assert.Equal(t, userLogin.Phone, phoneString)
	assert.Equal(t, err, nil)

}

func TestUserRegister(t *testing.T) {
	config.InitTest()

	repoUser := repository.NewUserRepository(config.DBTest)
	repoAddress := repository.NewAddressRepository(config.DBTest)
	repoDetail := repository.NewUserDetailRepository(config.DBTest)
	repository := repository.NewRepository(config.DBTest)

	service := service.NewUserService(repoUser, repoAddress, repoDetail, repository)

	phone := helpers.RandNumber(1000000000, 9999999999)
	phoneString := strconv.Itoa(phone)
	t.Log(phoneString)

	user, err := service.Register(request.UserRegister{
		Phone:       phoneString,
		Password:    "12345678",
		Name:        "test",
		Surname:     "test",
		Email:       "deneme@gmail.com",
		Birthday:    "01.01.2000",
		Districtfk:  1,
		Description: "test",
	})

	assert.Equal(t, user.Phone, phoneString)
	assert.Equal(t, err, nil)
	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Surname, "test")
	assert.Equal(t, user.Email, "deneme@gmail.com")
}

func TestUserUpdate(t *testing.T) {
	config.InitTest()

	repoUser := repository.NewUserRepository(config.DBTest)
	repoAddress := repository.NewAddressRepository(config.DBTest)
	repoDetail := repository.NewUserDetailRepository(config.DBTest)
	repository := repository.NewRepository(config.DBTest)

	service := service.NewUserService(repoUser, repoAddress, repoDetail, repository)

	phone := helpers.RandNumber(1111111111, 9999999999)
	phoneString := strconv.Itoa(phone)
	user, err := service.Register(request.UserRegister{
		Phone:       phoneString,
		Password:    "12345678",
		Name:        "test",
		Surname:     "test",
		Email:       "deneme@gmail.com",
		Birthday:    "01.01.2000",
		Districtfk:  1,
		Description: "test",
	})

	assert.Equal(t, user.Phone, phoneString)
	assert.Equal(t, err, nil)
	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Surname, "test")
	assert.Equal(t, user.Email, "deneme@gmail.com")

	err = service.Update(request.UserUpdate{
		ID:       user.ID,
		Name:     "test2",
		Surname:  "test2",
		Phone:    phoneString,
		Email:    "a@gmail.com",
		Birthday: "01.01.2001",
	})

	assert.Equal(t, err, nil)
}
