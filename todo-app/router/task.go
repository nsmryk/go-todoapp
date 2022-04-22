package router

import (
	"net/http"
	"todo-app/model"

	"github.com/google/uuid"
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

	return c.Render(http.StatusOK, "tasks.html", tasks)
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
	GetTasksHandler(c)
	return c.JSON(http.StatusOK, task)
}

func ChangeFinishiedTaskHandler(c echo.Context) error {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	err = model.ChangeFinishedTask(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	return c.NoContent(http.StatusOK)
}

func DeleteTaskHandler(c echo.Context) error {
	taskID, err := uuid.Parse(c.Param("taskID"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	err = model.DeleteTask(taskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.NoContent(http.StatusOK)
}
