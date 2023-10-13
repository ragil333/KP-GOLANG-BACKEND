package RoleController

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
	"kp-elibrary-golang/respositories/RoleRepo"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func Store(c echo.Context) error {
	var role models.Role
	if err := c.Bind(&role); err != nil {
		return helpers.ClientErrors(c, err.Error())
	}
	if err := helpers.Validation(role); err != nil {
		return helpers.ValidationErrors(c, err)
	}
	result, err := RoleRepo.Store(&role)
	if err != nil {
		return helpers.ClientErrors(c, err)
	}
	return helpers.CreateData(c, &result)
}
func Index(c echo.Context) error {
	result, err := RoleRepo.Index()
	if err != nil {
		return helpers.ClientErrors(c, err)
	}
	return helpers.FetchData(c, result)
}
func Show(c echo.Context) error {
	id := c.Param("id")
	result, err := RoleRepo.Show(id)
	if err != nil {
		return helpers.DataNotFound(c)
	}
	return helpers.FetchData(c, result)
}
func Destroy(c echo.Context) error {
	id := c.Param("id")
	if err := RoleRepo.Destroy(id); err != nil {
		return helpers.ClientErrors(c, err)
	}
	return helpers.DeleteData(c)
}
func Update(c echo.Context) error {
	var data models.Role
	id := c.Param("id")
	v := validator.New()
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := v.Struct(&data); err != nil {
		return helpers.ValidationErrors(c, err.Error())
	}
	result, err := RoleRepo.Update(&data, id)
	if err != nil {
		return helpers.ClientErrors(c, err)
	}
	return helpers.UpdateData(c, &result)
}
