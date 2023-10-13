package routes

import (
	"kp-elibrary-golang/controllers"

	"github.com/labstack/echo"
)

func RoleRoutes(e *echo.Group) {

	e.Add("GET", "role", controllers.AllRole)
	e.Add("GET", "role/:id", controllers.ShowRole)
	e.Add("POST", "role", controllers.StoreRole)
	e.Add("PUT", "role/:id", controllers.UpdateRole)
	e.Add("DELETE", "role/:id", controllers.DestroyRole)
}
