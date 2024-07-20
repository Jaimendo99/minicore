package project

import (
	"github.com/labstack/echo/v4"
	"minicore/inits"
	"minicore/models"
	"minicore/views/project"
	"net/http"
)

func GetProjects() echo.HandlerFunc {
	return func(c echo.Context) error {
		var projects []models.Project
		inits.DB.Find(&projects)

		if c.Request().Header.Get("Content-Type") == "application/json" {
			return c.JSON(http.StatusOK, projects)
		} else {
			return project.ProjectsView(projects).Render(c.Request().Context(), c.Response().Writer)
		}
	}
}
