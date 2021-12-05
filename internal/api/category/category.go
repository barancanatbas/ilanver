package category

import (
	"ilanver/internal/helpers"
	"ilanver/internal/models"
	"ilanver/repository"
	"ilanver/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Insert(c echo.Context) error {
	var req request.CategoryInsert
	if helpers.Validator(&c, &req) != nil {
		return nil
	}
	// main category kontrolü
	if req.MainCategory != 0 {
		count := repository.Get().Category().Exists(req.MainCategory)
		if !count {
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

func MainCategory(c echo.Context) error {

	categorys, err := repository.Get().Category().MainCategory()
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.Response(nil, "Bir hata oldu"))
	}

	return c.JSON(http.StatusBadRequest, helpers.Response(categorys, "Başarılı"))

}
