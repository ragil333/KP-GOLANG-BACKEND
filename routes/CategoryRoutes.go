package routes

import (
	"kp-elibrary-golang/controllers"

	"github.com/labstack/echo"
)

func CategoryRoutes(e *echo.Group) {
	e.Add("GET", "category", controllers.AllCategory)
	e.Add("GET", "category/:id", controllers.ShowCategory)
	e.Add("POST", "category", controllers.StoreCategory)
	e.Add("PUT", "category/:id", controllers.UpdateCategory)
	e.Add("DELETE", "category/:id", controllers.DestroyCategory)
}
