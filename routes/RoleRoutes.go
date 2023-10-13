package routes

import (
	"kp-elibrary-golang/controllers/RoleController"

	"github.com/labstack/echo"
)

func RoleRoutes(e *echo.Group) {

	e.Add("GET", "role", RoleController.Index)
	e.Add("GET", "role/:id", RoleController.Show)
	e.Add("POST", "role", RoleController.Store)
	e.Add("PUT", "role/:id", RoleController.Update)
	e.Add("DELETE", "role/:id", RoleController.Destroy)
}
