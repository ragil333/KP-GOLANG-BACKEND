package AuthRepo

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"

	"golang.org/x/crypto/bcrypt"
)

func Login(credentials *models.Login) (*models.User, error) {
	var checkUser models.User
	if err := helpers.DB.Where("email = ?", &credentials.Email).Find(&checkUser).Error; err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(credentials.Password)); err != nil {
		return nil, err
	}
	return &checkUser, nil
}
