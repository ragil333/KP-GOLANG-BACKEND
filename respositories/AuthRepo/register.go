package AuthRepo

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"

	"golang.org/x/crypto/bcrypt"
)

func Register(user models.User) (models.User, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)

	if err := helpers.DB.Create(&user); err.Error != nil {
		return user, err.Error
	}
	return user, nil
}
