package employee

import (
	"minicore/inits"
	"minicore/models"
	"minicore/views/employee"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetEmployees() echo.HandlerFunc {
	return func(c echo.Context) error {
		var employees []models.Employee
		inits.DB.Find(&employees)

		if c.Request().Header.Get("Content-Type") == "application/json" {
			return c.JSON(http.StatusOK, employees)
		} else {
			return employee.EmployeesView(employees).Render(c.Request().Context(), c.Response().Writer)
		}
	}
}
