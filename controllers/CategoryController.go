package controllers

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func StoreCategory(c echo.Context) error {

	var category models.Category
	if err := c.Bind(&category); err != nil {
		return helpers.ClientErrors(c, err.Error())
	}
	if err := helpers.Validation(category); err != nil {
		return helpers.ValidationErrors(c, err)
	}
	helpers.DB.Create(&category)
	return helpers.CreateData(c, &category)
}
func AllCategory(c echo.Context) error {
	var category []models.Category
	err := helpers.DB.Find(&category).Error
	if err != nil {
		return helpers.DataNotFound(c)
	}
	return helpers.FetchData(c, category)
}
func ShowCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	if err := helpers.DB.Where("category_id=?", id).First(&category).Error; err != nil {
		return helpers.DataNotFound(c)

	}
	return helpers.FetchData(c, category)
}
func DestroyCategory(c echo.Context) error {
	id := c.Param("id")
	var category models.Category
	err := helpers.DB.Where("category_id=?", id).First(&category).Error
	if err != nil {
		return helpers.DataNotFound(c)
	}
	helpers.DB.Delete(&category)
	return helpers.DeleteData(c)
}
func UpdateCategory(c echo.Context) error {
	var data models.Category
	id := c.Param("id")
	v := validator.New()
	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := v.Struct(&data); err != nil {
		return helpers.ValidationErrors(c, err.Error())
	}
	var category models.Category
	if err := helpers.DB.Where("category_id = ?", id).First(&category).Error; err != nil {
		return helpers.DataNotFound(c)
	}
	helpers.DB.Model(&category).Updates(&data)
	return helpers.UpdateData(c, &category)
}
