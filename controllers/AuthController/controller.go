package AuthController

import (
	"kp-elibrary-golang/helpers"
	"kp-elibrary-golang/models"
	"kp-elibrary-golang/respositories/AuthRepo"

	"github.com/labstack/echo"
)

func Register(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return helpers.ClientErrors(c, err.Error())
	}
	if err := helpers.Validation(user); err != nil {
		return helpers.ValidationErrors(c, err)
	}
	user, err := AuthRepo.Register(user)
	if err != nil {
		return helpers.ClientErrors(c, err)
	}
	return helpers.CreateData(c, user)

}
func Login(c echo.Context) error {
	var credentials models.Login
	if err := c.Bind(&credentials); err != nil {
		return helpers.ClientErrors(c, err.Error())
	}
	if err := helpers.Validation(&credentials); err != nil {
		return helpers.ValidationErrors(c, err)
	}
	user, err := AuthRepo.Login(&credentials)
	if err != nil {
		return helpers.UnAuthorizedUser(c)
	}
	token, err := helpers.GetToken(user)
	if err != nil {
		return helpers.ServerErrors(c, err)
	}
	return helpers.UserAuthenticated(c, &token)
}
