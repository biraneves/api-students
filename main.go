package main

import (
	"fmt"
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
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handlers
func getStudents(c echo.Context) error {
	return c.String(http.StatusOK, "List with all students")
}

func createStudent(c echo.Context) error {
	return c.String(http.StatusOK, "Create a student")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	strMsg := fmt.Sprintf("Info from student %s", id)
	return c.String(http.StatusOK, strMsg)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	strMsg := fmt.Sprintf("Update infos of student %s", id)
	return c.String(http.StatusOK, strMsg)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	strMsg := fmt.Sprintf("Delete student %s", id)
	return c.String(http.StatusOK, strMsg)
}
