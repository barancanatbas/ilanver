package servicetest

import (
	"ilanver/internal/model"
	"ilanver/internal/repository/mocks"
	"ilanver/internal/service"
	"ilanver/request"
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

	req := request.UserUpdate{
		ID:       1,
		Name:     "test2",
		Surname:  "test2",
		Phone:    "5551755445",
		Email:    "a@gmail.com",
		Birthday: "01.01.2001",
	}

	userMock := mocks.NewIUserRepo(t)

	userMock.On("Get", req.ID).Return(model.User{}, nil)
	userMock.On("Update", mock.Anything).Return(nil)

	service := service.NewUserService(userMock, nil, nil, nil)

	err := service.Update(req)

	assert.Equal(t, err, nil)
}
