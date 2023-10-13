package RoleRepo

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
)

func Index(role []models.Role) ([]models.Role, error) {
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
func Show() {

}
func Update() {

}
func Destroy() {

}
