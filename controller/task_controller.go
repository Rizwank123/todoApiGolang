package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/rizwank123/models"
	"github.com/rizwank123/service"

)

type TaskController struct {
	ts service.TaskService
}

func NewTaskController(contr service.TaskService) TaskController {
	return TaskController{ts: contr}
}

// CreatTask godoc
//
//	@Summary		Creat a task
//	@Description	Get a new Task
//	@Tags			Task
//	@Accept			json
//	@Produce		json
//	@Param			t	body		models.Task	true	"Task object to be created"
//	@Success		200	{object}	models.Task
//	@Router			/tasks [post]
func (tc *TaskController) CreateTask(c echo.Context) error {
	var task models.Task

	if err := c.Bind(&task); err != nil {
		fmt.Printf("%v", err)
		return c.JSON(http.StatusBadRequest, "Bindig faild")
	}
	err := tc.ts.CreateTask(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Row Not Inserted")
	}
	return c.JSON(http.StatusAccepted, task)
}

// DeleteTask godoc
//
//	@Summary		Delete task
//	@Description	delete a task by id
//	@Tags			Task
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Task ID"
//	@Success		200	{string}	deleted	successfully
//	@Router			/tasks/{id} [delete]
func (tc *TaskController) Delete(c echo.Context) error {
	id := c.Param("id")
	tid, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	tc.ts.Delete(tid)
	return c.JSON(http.StatusOK, "Record Deleted Successfully")
}

// Update godoc
//
//	@Summary		update  task
//	@Description	update  task by id
//	@Tags			Task
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string		true	"Task ID"
//	@Param			task	body		models.Task	true	"Task"
//	@Success		200		{object}	models.Task
//	@Router			/tasks/{id} [put]
func (tc *TaskController) Update(c echo.Context) error {
	var task models.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, "Bindig faild")
	}
	id := c.Param("id")
	tid, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	tsk, err := tc.ts.Update(tid, &task)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task Not foud")
	}
	return c.JSON(http.StatusOK, tsk)
}

// GetTask godoc
//
//	@Summary		Get a task
//	@Description	Get a task by id
//	@Tags			Task
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Task ID"
//	@Success		200	{object}	models.Task
//	@Router			/tasks/{id} [get]
func (tc *TaskController) FindByID(c echo.Context) error {
	id := c.Param("id")
	tid, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	tsk, err := tc.ts.FindByID(tid)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Task Not Found")
	}
	return c.JSON(http.StatusOK, tsk)
}

// ListTask godoc
//
//	@Summary		get all task
//	@Description	Get list of tasks
//	@Tags			Task
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Task
//	@Router			/tasks [get]
func (tc *TaskController) FindAllTask(c echo.Context) error {
	tasks, err := tc.ts.FindAllTask()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, tasks)
}

// ListTask godoc
//
//	@Summary		get all task
//	@Description	Get list of tasks
//	@Tags			Task
//	@Accept			json
//	@Produce		json
//	@Param			userId	path		string	true	"User ID"
//	@Success		200	{object}	models.Task
//	@Router			/tasks/users/{userId} [get]
func (tc *TaskController) FindByUserId(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		fmt.Printf("%d,%v\n", uid, err)
		return c.JSON(http.StatusNotFound, "Id Not Found")
	}
	task, err := tc.ts.FindByUserId(uid)
	if err != nil {
		return c.JSON(http.StatusNotFound, "user Not foud")
	}
	return c.JSON(http.StatusOK, task)
}
