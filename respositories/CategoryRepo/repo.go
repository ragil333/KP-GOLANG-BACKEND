package CategoryRepo

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
)

func Index() {

}

func Store(category *models.Category) (*models.Category, error) {
	result := helpers.DB.Create(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
func Show() {

}
func Update() {

}
func Destroy() {

}
