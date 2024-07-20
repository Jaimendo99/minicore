package UiCOntrollers

import (
	"github.com/labstack/echo/v4"
	"minicore/views/components"
)

func GetHomePage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return components.MainLayout("Home").Render(c.Request().Context(), c.Response().Writer)
	}
}
