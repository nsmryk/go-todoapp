package router

import (
	"net/http"
	"todo-app/model"

	"github.com/labstack/echo/v4"
)

type ReqTask struct {
	Name string `json:"name"`
}

func GetTasksHandler(c echo.Context) error {
	tasks, err := model.GetTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, tasks)
}

func AddTaskHandler(c echo.Context) error {
	var req ReqTask
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	var task *model.Task
	task, err = model.AddTask(req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.JSON(http.StatusOK, task)
}
