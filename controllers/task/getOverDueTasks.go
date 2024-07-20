package task

import (
	"github.com/labstack/echo/v4"
	"minicore/inits"
	"minicore/models"
	"minicore/views/task"
	"net/http"
	"time"
)

func GetOverDueTasks() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tasks []models.Task
		// Asegúrate de cargar las relaciones Employee y Project
		if err := inits.DB.Preload("Employee").Preload("Project").Find(&tasks).Error; err != nil {
			return err
		}

		var startDate = c.QueryParam("startDate")
		if startDate != "" {
			startDateTS, err := models.DateStringToInt32(startDate)
			if err != nil {
				return err
			}
			tasks = models.FilterTasksByStartDate(tasks, startDateTS)
		}

		var due []models.OverDue
		err := error(nil)

		var endDate = c.QueryParam("endDate")
		if endDate != "" {
			endDateTS, err := models.DateStringToInt32(endDate)
			due, err = models.TaskToOverDueList(tasks, time.Unix(int64(endDateTS), 0))
			if err != nil {
				return err
			}

		} else {
			due, err = models.TaskToOverDueList(tasks, time.Now())
			if err != nil {
				return err
			}
		}

		if c.Request().Header.Get("Content-Type") == "application/json" {
			return c.JSON(http.StatusOK, due)
		} else {
			if c.QueryParam("startDate") == "" {
				return task.OverDueTasksView(due).Render(c.Request().Context(), c.Response().Writer)
			}
			return task.OverDueTable(due).Render(c.Request().Context(), c.Response().Writer)
		}
	}
}
