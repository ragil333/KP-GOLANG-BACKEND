package RoleController

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
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
	helpers.DB.Create(&role)
	return helpers.CreateData(c, &role)
}
func Index(c echo.Context) error {
	var category []models.Role
	err := helpers.DB.Find(&category).Error
	if err != nil {
		return helpers.DataNotFound(c)
	}
	return helpers.FetchData(c, category)
}
func Show(c echo.Context) error {
	id := c.Param("id")
	var category models.Role
	if err := helpers.DB.Where("role_id=?", id).First(&category).Error; err != nil {
		return helpers.DataNotFound(c)

	}
	return helpers.FetchData(c, category)
}
func Destroy(c echo.Context) error {
	id := c.Param("id")
	var category models.Role
	err := helpers.DB.Where("role_id=?", id).First(&category).Error
	if err != nil {
		return helpers.DataNotFound(c)
	}
	helpers.DB.Delete(&category)
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
	var category models.Role
	if err := helpers.DB.Where("role_id = ?", id).First(&category).Error; err != nil {
		return helpers.DataNotFound(c)
	}
	helpers.DB.Model(&category).Updates(&data)
	return helpers.UpdateData(c, &category)
}