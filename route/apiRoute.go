package route

import (
	"github.com/labstack/echo/v4"

	"github.com/rizwank123/controller"
)

type Api struct {
	uc controller.UserController
	tc controller.TaskController
}

func NewApi(uc controller.UserController, tc controller.TaskController) *Api {
	return &Api{uc: uc, tc: tc}
}
func (a *Api) IntitRoute(e *echo.Echo) {
	e.GET("/users/:id", a.uc.FindByID)
	e.POST("/users", a.uc.CreateUser)
	e.PUT("users/:id", a.uc.Update)
	e.DELETE("/users/:id", a.uc.Delete)
	e.DELETE("/tasks/:id", a.tc.Delete)
	e.POST("/tasks", a.tc.CreateTask)
	e.PUT("/tasks/:id", a.tc.Update)
	e.GET("/tasks", a.tc.FindAllTask)
	e.GET("/tasks/:id", a.tc.FindByID)
	e.GET("/tasks/users/:userId", a.tc.FindByUserId)
}
