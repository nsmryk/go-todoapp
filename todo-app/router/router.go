package router

import (
	_ "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouter(e *echo.Echo) error {

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/tasks", GetTasksHandler)
	e.POST("/api/tasks", AddTaskHandler)
	e.PUT("/api/tasks/:taskID", ChangeFinishiedTaskHandler)
	e.DELETE("/api/tasks/:taskID", DeleteTaskHandler)

	err := e.Start(":8000")
	return err
}
