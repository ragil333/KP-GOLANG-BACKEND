package RoleRepo

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
)

func Index() ([]models.Role, error) {
	var role []models.Role
	result := helpers.DB.Preload("User").Find(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}
func Store(role *models.Role) (*models.Role, error) {
	result := helpers.DB.Create(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return role, nil
}
func Show(id any) (*models.Role, error) {
	var role models.Role
	result := helpers.DB.Where("role_id=?", id).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	return &role, nil
}
func Update(data *models.Role, id any) (*models.Role, error) {
	var role models.Role
	result := helpers.DB.Where("role_id = ?", id).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}
	helpers.DB.Model(&role).Updates(&data)
	return data, nil
}
func Destroy(id any) error {
	var role models.Role
	err := helpers.DB.Where("role_id=?", id).First(&role).Error
	if err != nil {
		return err
	}
	helpers.DB.Delete(&role)
	return nil
}
