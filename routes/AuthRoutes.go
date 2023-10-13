package routes

import (
	"kp-elibrary-golang/controllers/AuthController"

	"github.com/labstack/echo"
)

func AuthRoutes(e *echo.Group) {
	e.Add("POST", "auth/register", AuthController.Register)
	e.Add("POST", "auth/login", AuthController.Login)
}
