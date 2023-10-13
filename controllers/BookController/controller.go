package BookController

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
	"kp-elibrary-golang/respositories/BookRepo"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Store(c echo.Context) error {

	var book models.Book
	if err := c.Bind(&book); err != nil {
		return helpers.ClientErrors(c, err.Error())
	}
	if err := helpers.Validation(book); err != nil {
		return helpers.ValidationErrors(c, err)
	}

	result, err := BookRepo.Store(&book)
	if err != nil {
		return helpers.ClientErrors(c, err)
	}

	return helpers.CreateData(c, &result)
}
func Index(c echo.Context) error {
	result, err := BookRepo.Index()
	if err != nil {
		helpers.ClientErrors(c, err)
	}
	return helpers.FetchData(c, result)
}
func Show(c echo.Context) error {
	id := c.Param("id")
	result, err := BookRepo.Show(id)
	if err != nil {
		return helpers.ClientErrors(c, err)
	}
	return helpers.FetchData(c, result)
}
func Destroy(c echo.Context) error {
	id := c.Param("id")
	if err := BookRepo.Destroy(id); err != nil {
		return helpers.ClientErrors(c, err)
	}
	return helpers.DeleteData(c)
}
func Update(c echo.Context) error {
	var data models.Book
	id := c.Param("id")
	v := validator.New()
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := v.Struct(&data); err != nil {
		return helpers.ValidationErrors(c, err.Error())
	}
	result, err := BookRepo.Update(data, id)
	if err != nil {
		return helpers.ClientErrors(c, err)
	}
	return helpers.UpdateData(c, &result)
}
