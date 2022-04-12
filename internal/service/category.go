package service

import (
	"ilanver/internal/model"
	"ilanver/internal/repository"
	"ilanver/request"
	"strconv"
)

type ICategoryService interface {
	GetAll(page string) (interface{}, error)
	GetSubCategories(id string) ([]model.Category, error)
	Insert(req request.InsertCategory) error
	Update(req request.UpdateCategory) error
	Delete(id string) error
}

type CategoryService struct {
	repoCategory repository.ICategoryRepository
	repository   repository.IRepository
}

func NewCategoryService(repoCategory repository.ICategoryRepository, repository repository.IRepository) ICategoryService {
	return CategoryService{
		repoCategory: repoCategory,
		repository:   repository,
	}
}

func (c CategoryService) GetAll(page string) (interface{}, error) {

	return c.repoCategory.GetAll()
}

func (c CategoryService) GetSubCategories(id string) ([]model.Category, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return []model.Category{}, err
	}
	return c.repoCategory.GetSubCategories(uint(idInt))
}

func (c CategoryService) Insert(req request.InsertCategory) error {
	data := model.Category{
		CategoryName: req.Name,
		MainCategory: req.MainCategory,
	}

	err := c.repoCategory.Insert(&data)

	return err
}

func (c CategoryService) Update(req request.UpdateCategory) error {

	category, err := c.repoCategory.GetByID(int(req.ID))

	category.CategoryName = req.Name
	category.MainCategory = req.MainCategory

	err = c.repoCategory.Update(category)

	return err
}

func (c CategoryService) Delete(id string) error {
	// TODO: burada silme işlemi yapılacak fakat alt kategorilerin silinmesi gerekiyor.

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	ids := findDeletedCategories(c.repoCategory, uint(idInt))
	ids = append(ids, idInt)

	err = c.repoCategory.DeleteWitchInQuery(ids)
	return err
}

func findDeletedCategories(repo repository.ICategoryRepository, id uint) []int {
	var ids []int
	categories, _ := repo.GetSubCategories(id)

	if len(categories) <= 0 {
		return ids
	}

	for _, category := range categories {
		ids = append(ids, int(category.ID))
		ids = append(ids, findDeletedCategories(repo, category.ID)...)
	}

	return ids
}
