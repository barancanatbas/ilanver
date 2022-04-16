package repositorytest

import (
	"ilanver/internal/config"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInsertDetail(t *testing.T) {
	config.InitTest()
	repoDetail := repository.NewUserDetailRepository(config.DBTest)

	repoAddress := repository.NewAddressRepository(config.DBTest)

	// save address
	address := model.Adress{
		Districtfk: 1,
		Detail:     "test address",
	}

	err := repoAddress.Save(&address)

	assert.Equal(t, err, nil)

	detail := model.UserDetail{
		ProfilePhoto: "test profile photo",
		Adressfk:     address.ID,
	}

	err = repoDetail.Save(&detail)

	assert.Equal(t, detail.Adressfk, address.ID)
	assert.Equal(t, err, nil)

}

func TestGetByIdDetail(t *testing.T) {
	config.InitTest()
	repoDetail := repository.NewUserDetailRepository(config.DBTest)

	repoAddress := repository.NewAddressRepository(config.DBTest)

	// save address
	address := model.Adress{
		Districtfk: 1,
		Detail:     "test address",
	}

	err := repoAddress.Save(&address)

	assert.Equal(t, err, nil)

	detail := model.UserDetail{
		ProfilePhoto: "test profile photo",
		Adressfk:     address.ID,
	}

	err = repoDetail.Save(&detail)

	assert.Equal(t, detail.Adressfk, address.ID)
	assert.Equal(t, err, nil)

	getDetail, err := repoDetail.GetByID(detail.ID)

	assert.Equal(t, err, nil)
	assert.Equal(t, getDetail.ID, detail.ID)
	assert.Equal(t, getDetail.Adressfk, detail.Adressfk)
}
