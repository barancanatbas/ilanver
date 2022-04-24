package servicetest

import (
	"ilanver/internal/config"
	"ilanver/internal/helpers"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"ilanver/internal/repository/mocks"
	"ilanver/internal/service"
	"ilanver/request"
	"strconv"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestUserLogin(t *testing.T) {
	Usermock := mocks.NewIUserRepo(t)

	reqUser := request.UserLogin{
		Phone:    "1111111111",
		Password: "12345678",
	}

	Usermock.On("Login", reqUser.Phone).Return(model.User{
		Phone:    reqUser.Phone,
		Password: "$2a$04$8prBxe1ZupRaRpRzMgnWuOpZ9HTBEj2IuBtLRLrswTH9jLt5iu9bS",
	}, nil)

	userService := service.NewUserService(Usermock, nil, nil, nil)

	user, _, err := userService.Login(reqUser)

	assert.Equal(t, user.Phone, reqUser.Phone)
	assert.Equal(t, err, nil)

}

func TestUserRegister(t *testing.T) {
	var db *gorm.DB
	userMock := mocks.NewIUserRepo(t)
	addressMock := mocks.NewIAddressRepo(t)
	detailMock := mocks.NewIUserDetailRepo(t)
	repositoryMock := mocks.NewIRepository(t)

	request := request.UserRegister{
		Name:        "test",
		Surname:     "test",
		Phone:       "1111111111",
		Password:    "12345678",
		Email:       "",
		Birthday:    "01.01.2000",
		Districtfk:  1,
		Description: "test",
	}

	// repository mock test case
	repositoryMock.On("CreateTX").Return(db)
	repositoryMock.On("Commit").Return(nil)

	addressMock.On("WithTx", db).Return(addressMock)
	addressMock.On("Save", mock.Anything).Return(nil)

	detailMock.On("WithTx", db).Return(detailMock)
	detailMock.On("Save", mock.Anything).Return(nil)

	userMock.On("WithTx", db).Return(userMock)
	userMock.On("Save", mock.Anything).Return(nil)

	service := service.NewUserService(userMock, addressMock, detailMock, repositoryMock)

	_, err := service.Register(request)

	assert.Equal(t, err, nil)

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
