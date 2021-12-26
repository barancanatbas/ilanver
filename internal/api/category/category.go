package category

import (
	"fmt"
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

	category := models.Category{}
	// main category kontrolü
	if req.MainCategory > 0 {
		_, err, rowCount := repository.Get().Category().Exists(req.MainCategory)
		if err != nil || rowCount == 0 {
			return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Main kategory bulunamadı"))
		}
		category.MainCategory = req.MainCategory
	}

	// new category
	category.CategoryName = req.CategoryName

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
	category, err, rowCount := repository.Get().Category().Exists(req.Id)
	if err != nil || rowCount == 0 {
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

// @Summary Main Categories
// @Description root kategorileri getirir.
// @Tags category
// @Router /category/main [get]
func MainCategory(c echo.Context) error {
	categorys, err := repository.Get().Category().MainCategory()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Bir hata oldu"))
	}

	return c.JSON(http.StatusBadRequest, helpers.Response(categorys, "Başarılı"))
}

// @Summary Delete Category
// @Description Var olan kategori bilgilerini siler.
// @Tags category
// @Param body body request.CategoryDelete false " "
// @Router /category [Delete]
func Delete(c echo.Context) error {
	var req request.CategoryDelete
	if helpers.Validator(&c, &req) != nil {
		return nil
	}

	// category
	category, err, rowCount := repository.Get().Category().Exists(req.Id)
	if err != nil || rowCount == 0 {
		fmt.Println(rowCount)
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "kategori bulunamadı"))
	}

	err = repository.Get().Category().Delete(category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "kategori silinemedi"))
	}

	return c.JSON(http.StatusBadRequest, helpers.Response(nil, "kategori bulunamadı"))
}
