package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handlers
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List with all students")
}
