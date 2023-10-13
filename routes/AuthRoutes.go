package routes

import (
	"kp-elibrary-golang/controllers"

	"github.com/labstack/echo"
)

func AuthRoutes(e *echo.Group) {
	e.Add("POST", "auth/register", controllers.Register)
	e.Add("POST", "auth/login", controllers.Login)
}
