package router

import (
	_ "net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type Template struct {
// 	templates *template.Template
// }

// func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
// 	return t.templates.ExecuteTemplate(w, name, data)
// }

func SetRouter(e *echo.Echo) error {
	// t := &Template{
	// 	templates: template.Must(template.ParseGlob("template/*.html")),
	// }
	// e.Renderer = t

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/api/tasks", GetTasksHandler)
	e.POST("/api/tasks", AddTaskHandler)
	e.PUT("/api/tasks/:taskID", ChangeFinishiedTaskHandler)
	e.DELETE("/api/tasks/:taskID", DeleteTaskHandler)

	err := e.Start(":8000")
	return err
}
