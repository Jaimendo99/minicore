package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"minicore/controllers/UiCOntrollers"
	"minicore/controllers/employee"
	"minicore/controllers/project"
	"minicore/controllers/task"
	"minicore/inits"
	"os"
)

func init() {
	inits.LoadEnvVar()

	inits.InitDB()

	err := inits.Migrate(inits.DB)
	if err != nil {
		fmt.Println("error migrating db" + err.Error())
		panic(err)
	}
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", UiCOntrollers.GetHomePage())

	e.GET("/employees", employee.GetEmployees())
	e.GET("/projects", project.GetProjects())
	e.GET("/tasks", task.GetTasks())

	e.GET("/overdue", task.GetOverDueTasks())

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
