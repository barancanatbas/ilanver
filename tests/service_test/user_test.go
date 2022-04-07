package servicetest

import (
	"ilanver/internal/config"
	"ilanver/internal/repository"
	"ilanver/internal/service"
	"ilanver/request"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUserLogin(t *testing.T) {
	config.Init()

	repoUser := repository.NewUserRepository(config.DB)
	repoAddress := repository.NewAddressRepository(config.DB)
	repoDetail := repository.NewUserDetailRepository(config.DB)
	repository := repository.NewRepository(config.DB)

	service := service.NewUserService(repoUser, repoAddress, repoDetail, repository)

	user, _, err := service.Login(request.UserLogin{
		Phone:    "5555555555",
		Password: "12345678",
	})

	assert.Equal(t, user.Phone, "5555555555")
	assert.Equal(t, err, nil)

}

func TestUserRegister(t *testing.T) {
	config.Init()

	repoUser := repository.NewUserRepository(config.DB)
	repoAddress := repository.NewAddressRepository(config.DB)
	repoDetail := repository.NewUserDetailRepository(config.DB)
	repository := repository.NewRepository(config.DB)

	service := service.NewUserService(repoUser, repoAddress, repoDetail, repository)

	user, err := service.Register(request.UserRegister{
		Phone:       "1555555555",
		Password:    "12345678",
		Name:        "test",
		Surname:     "test",
		Email:       "deneme@gmail.com",
		Birthday:    "01.01.2000",
		Districtfk:  1,
		Description: "test",
	})

	assert.Equal(t, user.Phone, "1555555555")
	assert.Equal(t, err, nil)
	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Surname, "test")
	assert.Equal(t, user.Email, "deneme@gmail.com")
}

func TestUserUpdate(t *testing.T) {
	config.Init()

	repoUser := repository.NewUserRepository(config.DB)
	repoAddress := repository.NewAddressRepository(config.DB)
	repoDetail := repository.NewUserDetailRepository(config.DB)
	repository := repository.NewRepository(config.DB)

	service := service.NewUserService(repoUser, repoAddress, repoDetail, repository)

	user, err := service.Register(request.UserRegister{
		Phone:       "1155555555",
		Password:    "12345678",
		Name:        "test",
		Surname:     "test",
		Email:       "deneme@gmail.com",
		Birthday:    "01.01.2000",
		Districtfk:  1,
		Description: "test",
	})

	assert.Equal(t, user.Phone, "1155555555")
	assert.Equal(t, err, nil)
	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Surname, "test")
	assert.Equal(t, user.Email, "deneme@gmail.com")

	err = service.Update(request.UserUpdate{
		ID:       user.ID,
		Name:     "test2",
		Surname:  "test2",
		Email:    "a@gmail.com",
		Birthday: "01.01.2001",
	})

	assert.Equal(t, err, nil)
}
