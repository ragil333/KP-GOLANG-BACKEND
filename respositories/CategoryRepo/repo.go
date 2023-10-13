package CategoryRepo

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
)

func Index() ([]models.Category, error) {
	var category []models.Category
	result := helpers.DB.Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func Store(category *models.Category) (*models.Category, error) {
	result := helpers.DB.Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
func Show(id any) (*models.Category, error) {
	var category models.Category
	result := helpers.DB.Where("category_id=?", id).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}
func Update(data models.Category, id any) (*models.Category, error) {
	var category models.Category
	result := helpers.DB.Where("category_id = ?", id).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	helpers.DB.Model(&category).Updates(&data)
	return &category, nil
}
func Destroy(id any) error {
	var category models.Category
	result := helpers.DB.Where("category_id=?", id).First(&category)
	if result.Error != nil {
		return result.Error
	}
	helpers.DB.Delete(&category)
	return nil
}
