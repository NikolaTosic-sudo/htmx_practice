package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	page := newPage()
	e.Renderer = newTemplate()

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", page)
	})

	e.GET("/blocks", BlocksPage)

	e.GET("/contact/edit/:id", page.GetEditContact())

	e.POST("/contacts", page.AddNewContact())

	e.DELETE("/contacts/:id", page.DeleteContact())

	e.PUT("/contact/update/:id", page.UpdateContact())

	e.Logger.Fatal(e.Start(":8080"))
}
