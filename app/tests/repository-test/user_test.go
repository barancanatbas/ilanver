package repositorytest

import (
	"ilanver/internal/cache"
	"ilanver/internal/config"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"golang.org/x/crypto/bcrypt"
)

func TestInsertUser(t *testing.T) {

	config.InitTest()

	repoAddress := repository.NewAddressRepository(config.DBTest)
	repoDetail := repository.NewUserDetailRepository(config.DBTest)
	repoUser := repository.NewUserRepository(config.DBTest)

	// save address
	address := model.Adress{
		Districtfk: 1,
		Detail:     "test address",
	}

	err := repoAddress.Save(&address)

	assert.Equal(t, err, nil)

	// save detail
	detail := model.UserDetail{
		ProfilePhoto: "test profile photo",
		Adressfk:     address.ID,
	}

	err = repoDetail.Save(&detail)

	assert.Equal(t, detail.Adressfk, address.ID)
	assert.Equal(t, err, nil)

	user := model.User{
		Name:         "test name",
		Surname:      "test surname",
		Phone:        "5551551",
		Password:     "test",
		Email:        "baran@gmail.com",
		UserDetailfk: detail.ID,
		Birthday:     time.Now(),
	}

	err = repoUser.Save(&user)

	assert.Equal(t, err, nil)
	assert.Equal(t, user.UserDetailfk, detail.ID)
}

func TestLoginUser(t *testing.T) {

	config.InitTest()

	repoAddress := repository.NewAddressRepository(config.DBTest)
	repoDetail := repository.NewUserDetailRepository(config.DBTest)
	repoUser := repository.NewUserRepository(config.DBTest)

	// save address
	address := model.Adress{
		Districtfk: 1,
		Detail:     "test address",
	}

	err := repoAddress.Save(&address)

	assert.Equal(t, err, nil)

	// save detail
	detail := model.UserDetail{
		ProfilePhoto: "test profile photo",
		Adressfk:     address.ID,
	}

	err = repoDetail.Save(&detail)

	assert.Equal(t, detail.Adressfk, address.ID)
	assert.Equal(t, err, nil)

	password, _ := bcrypt.GenerateFromPassword([]byte("deneme"), 4)

	user := model.User{
		Name:         "test name",
		Surname:      "test surname",
		Phone:        "5555551",
		Password:     string(password),
		Email:        "baran@gmail.com",
		UserDetailfk: detail.ID,
		Birthday:     time.Now(),
	}

	err = repoUser.Save(&user)

	assert.Equal(t, err, nil)
	assert.Equal(t, user.UserDetailfk, detail.ID)

	loginUser, err := repoUser.Login(user.Phone)

	assert.Equal(t, err, nil)
	assert.Equal(t, loginUser.Password, string(password))
}

func TestGetByIDUser(t *testing.T) {

	config.InitTest()

	repoAddress := repository.NewAddressRepository(config.DBTest)
	repoDetail := repository.NewUserDetailRepository(config.DBTest)
	repoUser := repository.NewUserRepository(config.DBTest)

	// save address
	address := model.Adress{
		Districtfk: 1,
		Detail:     "test address",
	}

	err := repoAddress.Save(&address)

	assert.Equal(t, err, nil)

	// save detail
	detail := model.UserDetail{
		ProfilePhoto: "test profile photo",
		Adressfk:     address.ID,
	}

	err = repoDetail.Save(&detail)

	assert.Equal(t, detail.Adressfk, address.ID)
	assert.Equal(t, err, nil)

	user := model.User{
		Name:         "test name2",
		Surname:      "test surname2",
		Phone:        "5555552",
		Password:     "test2",
		Email:        "baran@gmail.com",
		UserDetailfk: detail.ID,
		Birthday:     time.Now(),
	}

	err = repoUser.Save(&user)

	assert.Equal(t, err, nil)
	assert.Equal(t, user.UserDetailfk, detail.ID)

	getUser, err := repoUser.Get(user.ID)

	assert.Equal(t, err, nil)
	assert.Equal(t, getUser.ID, user.ID)
}

func TestChangePasswordCode(t *testing.T) {

	config.Pool = config.NewPool()

	cache.SetHashCache("lostPassword:55555555", map[string]string{
		"code":    "1212",
		"phone":   "55555555",
		"ip":      ":1",
		"confirm": "false",
	})

	if !cache.Exists("lostPassword:55555555") {
		t.Error("cache not exists")
	}

	cacheData := cache.GetHashCache("lostPassword:55555555")

	assert.Equal(t, cacheData["code"], "1212")
	assert.Equal(t, cacheData["phone"], "55555555")
}
