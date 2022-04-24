package servicetest

import (
	"ilanver/internal/model"
	"ilanver/internal/repository/mocks"
	"ilanver/internal/service"
	"ilanver/request"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetAllCategory(t *testing.T) {
	mockCategory := mocks.NewICategoryRepository(t)

	categories := []model.Category{
		{
			CategoryName: "test category 1",
			MainCategory: 0,
		},
		{
			CategoryName: "test category 2",
			MainCategory: 0,
		},
		{
			CategoryName: "test category 3",
			MainCategory: 0,
		},
	}

	mockCategory.On("GetAll").Return(categories, nil)

	mockRepository := mocks.NewIRepository(t)

	service := service.NewCategoryService(mockCategory, mockRepository)

	want, err := service.GetAll()

	assert.Equal(t, err, nil)

	assert.Equal(t, want, categories)
}

func TestCategoryInsert(t *testing.T) {
	// burada bir mock repository oluşturduk.
	mockCategory := mocks.NewICategoryRepository(t)

	categories := request.InsertCategory{
		Name:         "test category",
		MainCategory: 0,
	}

	// burada repository katmanındaki işlemin olduğu yere hangi verilerin geleceğini ve hangi verinin döneceğini söylüyoruz.
	mockCategory.On("Insert", &model.Category{CategoryName: categories.Name, MainCategory: categories.MainCategory}).Return(nil)

	mockRepository := mocks.NewIRepository(t)

	service := service.NewCategoryService(mockCategory, mockRepository)

	// verileri gönderiyoruz ve işlemi gerçekleştiriyoruz.
	err := service.Insert(categories)

	assert.Equal(t, err, nil)
}

func TestCategoryUpdate(t *testing.T) {
	// burada bir mock repository oluşturduk.
	mockCategory := mocks.NewICategoryRepository(t)

	categories := request.UpdateCategory{
		ID:           1,
		Name:         "test category",
		MainCategory: 0,
	}

	mockCategory.On("Update", model.Category{CategoryName: categories.Name, MainCategory: categories.MainCategory}).Return(nil)

	mockCategory.On("GetByID", 1).Return(model.Category{CategoryName: categories.Name, MainCategory: categories.MainCategory}, nil)

	mockRepository := mocks.NewIRepository(t)

	service := service.NewCategoryService(mockCategory, mockRepository)

	err := service.Update(categories)

	assert.Equal(t, err, nil)
}

func TestCategoryDelete(t *testing.T) {
	mockCategory := mocks.NewICategoryRepository(t)

	mockCategory.On("DeleteWithInQuery", []int{2, 3, 4, 1}).Return(nil)

	mockCategory.On("GetSubCategories", uint(1)).Return([]model.Category{
		{
			CategoryName: "test category 1",
			MainCategory: 1,
			ID:           2,
		},
		{
			CategoryName: "test category 2",
			MainCategory: 1,
			ID:           3,
		},
		{
			CategoryName: "test category 3",
			MainCategory: 1,
			ID:           4,
		},
	}, nil)

	mockCategory.On("GetSubCategories", uint(2)).Return([]model.Category{}, nil)
	mockCategory.On("GetSubCategories", uint(3)).Return([]model.Category{}, nil)
	mockCategory.On("GetSubCategories", uint(4)).Return([]model.Category{}, nil)

	mockRepository := mocks.NewIRepository(t)

	service := service.NewCategoryService(mockCategory, mockRepository)

	err := service.Delete("1")

	assert.Equal(t, err, nil)
}

func TestCategoryGetSubCategories(t *testing.T) {
	mockCategory := mocks.NewICategoryRepository(t)

	mockCategory.On("GetSubCategories", uint(1)).Return([]model.Category{
		{
			CategoryName: "test category 1",
			MainCategory: 1,
			ID:           2,
		},
		{
			CategoryName: "test category 2",
			MainCategory: 1,
			ID:           3,
		},
		{
			CategoryName: "test category 3",
			MainCategory: 1,
			ID:           4,
		},
	}, nil)

	mockRepository := mocks.NewIRepository(t)

	service := service.NewCategoryService(mockCategory, mockRepository)

	subCategories, err := service.GetSubCategories("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(subCategories), 3)
}
