package repositorytest

import (
	"ilanver/internal/config"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPstateGetAll(t *testing.T) {
	config.Init()
	repo := repository.NewProductStateRepository(config.DB)

	_, err := repo.GetAll()

	assert.Equal(t, nil, err)
}

func TestInsertPstate(t *testing.T) {
	config.Init()
	repo := repository.NewProductStateRepository(config.DB)

	pstate := model.ProductState{
		State: "test product state",
	}

	err := repo.Insert(&pstate)

	assert.Equal(t, nil, err)
}

func TestGetByIDPstate(t *testing.T) {
	config.Init()
	repo := repository.NewProductStateRepository(config.DB)

	pstate := model.ProductState{
		State: "test product state",
	}

	err := repo.Insert(&pstate)

	assert.Equal(t, nil, err)

	pstateNew, err := repo.GetByID(int(pstate.ID))

	assert.Equal(t, nil, err)
	assert.Equal(t, pstateNew.State, pstate.State)
	assert.Equal(t, pstateNew.ID, pstate.ID)
}

func TestUpdatePstate(t *testing.T) {
	config.Init()
	repo := repository.NewProductStateRepository(config.DB)

	pstate := model.ProductState{
		State: "test product state",
	}

	err := repo.Insert(&pstate)

	assert.Equal(t, nil, err)

	pstate.State = "test product state updated"

	err = repo.Update(&pstate)

	assert.Equal(t, nil, err)

	pstateNew, err := repo.GetByID(int(pstate.ID))

	assert.Equal(t, nil, err)
	assert.Equal(t, pstateNew.State, pstate.State)
	assert.Equal(t, pstateNew.ID, pstate.ID)
}

func TestDeletePstate(t *testing.T) {
	config.Init()
	repo := repository.NewProductStateRepository(config.DB)

	pstate := model.ProductState{
		State: "test product state",
	}

	err := repo.Insert(&pstate)

	assert.Equal(t, nil, err)

	err = repo.Delete(pstate.ID)

	assert.Equal(t, nil, err)
}
