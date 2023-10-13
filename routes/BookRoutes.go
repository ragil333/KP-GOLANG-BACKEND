package routes

import (
	"kp-elibrary-golang/controllers/BookController"

	"github.com/labstack/echo"
)

func BookRoutes(e *echo.Group) {
	e.Add("GET", "book", BookController.Index)
	e.Add("GET", "book/:id", BookController.Show)
	e.Add("POST", "book", BookController.Store)
	e.Add("PUT", "book/:id", BookController.Update)
	e.Add("DELETE", "book/:id", BookController.Destroy)
}
