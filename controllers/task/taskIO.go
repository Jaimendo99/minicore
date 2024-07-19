package task

import (
	"minicore/inits"
	"minicore/models"
	"minicore/views/task"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tasks []models.Task
		inits.DB.Find(&tasks)

		if c.Request().Header.Get("Content-Type") == "application/json" {
			return c.JSON(http.StatusOK, tasks)
		} else {
			return task.TasksView(tasks).Render(c.Request().Context(), c.Response().Writer)
		}
	}
}
