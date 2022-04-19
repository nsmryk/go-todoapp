package router

import (
	"net/http"
	"todo-app/model"

	"github.com/labstack/echo/v4"
)

func GetTasksHandler(c echo.Context) error {
	tasks, err := model.GetTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, tasks)
}
