package repositorytest

import (
	"ilanver/internal/config"
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"strconv"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestInsertMainCategory(t *testing.T) {
	config.Init()
	repo := repository.NewCategoryRepository(config.DB)

	category := model.Category{
		CategoryName: "test category",
		MainCategory: 0,
	}

	err := repo.Insert(&category)

	assert.Equal(t, err, nil)
}

func TestInsertChildCategory(t *testing.T) {
	config.Init()
	repo := repository.NewCategoryRepository(config.DB)

	category := model.Category{
		CategoryName: "test category",
		MainCategory: 0,
	}

	err := repo.Insert(&category)

	assert.Equal(t, err, nil)

	category2 := model.Category{
		CategoryName: "test category",
		MainCategory: category.ID,
	}

	err = repo.Insert(&category2)

	assert.Equal(t, err, nil)
}

func TestGetMainCategories(t *testing.T) {
	config.Init()
	repo := repository.NewCategoryRepository(config.DB)

	_, err := repo.GetAll()

	assert.Equal(t, err, nil)

}

func TestGetSubCategories(t *testing.T) {
	config.Init()
	repo := repository.NewCategoryRepository(config.DB)

	category := model.Category{
		CategoryName: "test category",
		MainCategory: 0,
	}

	err := repo.Insert(&category)

	assert.Equal(t, err, nil)

	for i := 0; i < 3; i++ {

		category2 := model.Category{
			CategoryName: "test category : " + strconv.Itoa(i),
			MainCategory: category.ID,
		}

		err = repo.Insert(&category2)

		assert.Equal(t, err, nil)
	}

	categories, err := repo.GetSubCategories(category.ID)

	assert.Equal(t, err, nil)

	assert.Equal(t, len(categories), 3)
}

func TestUpdateCategory(t *testing.T) {
	config.Init()
	repo := repository.NewCategoryRepository(config.DB)

	category := model.Category{
		CategoryName: "test category",
		MainCategory: 0,
	}

	err := repo.Insert(&category)

	assert.Equal(t, err, nil)

	category.CategoryName = "test category updated"

	err = repo.Update(category)

	assert.Equal(t, err, nil)
}
