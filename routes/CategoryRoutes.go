package routes

import (
	"kp-elibrary-golang/controllers/CategoryController"

	"github.com/labstack/echo"
)

func CategoryRoutes(e *echo.Group) {
	e.Add("GET", "category", CategoryController.Index)
	e.Add("GET", "category/:id", CategoryController.Show)
	e.Add("POST", "category", CategoryController.Store)
	e.Add("PUT", "category/:id", CategoryController.Update)
	e.Add("DELETE", "category/:id", CategoryController.Destroy)
}
