package category

import (
	"ilanver/internal/helpers"
	"ilanver/internal/models"
	"ilanver/repository"
	"ilanver/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Insert Category
// @Description yeni kategory ekler
// @Tags category
// @Param body body request.CategoryInsert false " "
// @Router /category [post]
func Insert(c echo.Context) error {
	var req request.CategoryInsert
	if helpers.Validator(&c, &req) != nil {
		return nil
	}
	// main category kontrolü
	if req.MainCategory != 0 {
		_, err := repository.Get().Category().Exists(req.MainCategory)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Main kategory bulunamadı"))
		}
	}

	// new category
	category := models.Category{
		CategoryName: req.CategoryName,
		MainCategory: req.MainCategory,
		Src:          "",
	}

	err := repository.Get().Category().Insert(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Kategori kayıt edilemedi"))
	}

	return c.JSON(http.StatusBadRequest, helpers.Response(req, "Kategory eklendi"))
}

// @Summary update Category
// @Description var olan kategoriyi günceller
// @Tags category
// @Param body body request.CategoryUpdate false " "
// @Router /category [put]
func Update(c echo.Context) error {
	var req request.CategoryUpdate
	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	// exists category
	category, err := repository.Get().Category().Exists(req.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "kategori bulunamadı."))
	}

	// prepare update category
	category.CategoryName = req.CategoryName
	if req.MainCategory > 0 {

		// exists main category
		mainCategory, err := repository.Get().Category().ExistsMain(req.MainCategory)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helpers.Response(nil, "ana kategori bulunamadı."))
		}

		category.MainCategory = mainCategory.ID
	}

	err = repository.Get().Category().Update(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "kategori güncellenemedi."))
	}
	return c.JSON(http.StatusBadRequest, helpers.Response(nil, "kategori başarıyla güncellendi."))
}

func MainCategory(c echo.Context) error {

	categorys, err := repository.Get().Category().MainCategory()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Bir hata oldu"))
	}

	return c.JSON(http.StatusBadRequest, helpers.Response(categorys, "Başarılı"))

}
