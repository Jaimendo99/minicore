package project

import (
	"github.com/labstack/echo/v4"
	"minicore/inits"
	"minicore/models"
	"net/http"
)

func GetProjects() echo.HandlerFunc {
	return func(c echo.Context) error {
		var projects []models.Project
		inits.DB.Find(&projects)
		return c.JSON(http.StatusOK, projects)
	}
}
