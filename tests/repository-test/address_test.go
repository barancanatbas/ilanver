package repositorytest

import (
	"ilanver/internal/config"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInsertAddress(t *testing.T) {
	config.Init()
	repo := repository.NewAddressRepository(config.DB)

	address := model.Adress{
		Districtfk: 1,
		Detail:     "test address",
	}

	err := repo.Save(&address)

	assert.Equal(t, err, nil)
}

func TestGetById(t *testing.T) {
	config.Init()
	repo := repository.NewAddressRepository(config.DB)

	address := model.Adress{
		Districtfk: 1,
		Detail:     "test address",
	}

	err := repo.Save(&address)

	assert.Equal(t, err, nil)

	getAddress, err := repo.GetByID(address.ID)

	assert.Equal(t, err, nil)
	assert.Equal(t, getAddress.Detail, address.Detail)
}

func TestUpdateAddress(t *testing.T) {
	config.Init()
	repo := repository.NewAddressRepository(config.DB)

	address := model.Adress{
		Districtfk: 1,
		Detail:     "test address",
	}

	err := repo.Save(&address)

	assert.Equal(t, err, nil)

	address.Detail = "test address 2"

	err = repo.Update(&address)

	assert.Equal(t, err, nil)

	getAddress, err := repo.GetByID(address.ID)

	assert.Equal(t, err, nil)
	assert.Equal(t, getAddress.Detail, address.Detail)
}
