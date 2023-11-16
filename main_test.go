package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rizwank123/controller"
	"github.com/rizwank123/models"
)

type TaskServiceMock struct {
	mock.Mock
}

func (m *TaskServiceMock) CreateTask(tsk *models.Task) error {
	args := m.Called(tsk)
	return args.Error(1)
}
func (m *TaskServiceMock) Delete(id int) error {
	args := m.Called()
	return args.Error(1)
}
func (m *TaskServiceMock) FindAllTask() ([]models.Task, error) {
	args := m.Called()
	return args.Get(0).([]models.Task), args.Error(1)
}
func (m *TaskServiceMock) FindByID(id int) (models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(models.Task), args.Error(1)
}
func (m *TaskServiceMock) FindByUserId(useId int) ([]models.Task, error) {
	args := m.Called(useId)
	return args.Get(0).([]models.Task), args.Error(1)
}
func (m *TaskServiceMock) Update(id int, task *models.Task) (models.Task, error) {
	args := m.Called(id, task)
	return args.Get(0).(models.Task), args.Error(1)
}

func TestCreateTask(t *testing.T) {
	// Initialize mock service
	mockService := &TaskServiceMock{}

	// Create mock controller
	mockController := controller.NewTaskController(mockService)

	// Define a sample task payload
	payload := `{
		"desc": "learning go test",
		"id": 1,
		"name": "Learning unit testing",
		"status": "Started",
		"user_id": 1
	}`

	// Create a request with the defined payload
	req := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBufferString(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// Create a response recorder and context
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// Define the expected created task based on the payload
	expectedTask := &models.Task{
		ID:     1,
		Name:   "Learning unit testing",
		Desc:   "learning go test",
		UserId: 1,
		Status: "Started",
	}

	// Mock the CreateTask method of the TaskService
	mockService.On("CreateTask", mock.AnythingOfType("*models.Task")).Return(expectedTask, nil)

	// Call the CreateTask function of the controller
	err := mockController.CreateTask(c)

	// Perform assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusAccepted, rec.Code)
	mockService.AssertCalled(t, "CreateTask", mock.AnythingOfType("*models.Task"))
}

func TestDelete(t *testing.T) {
	// Initialize mock service
	mockService := &TaskServiceMock{}

	// Create mock controller or initialize
	mockController := controller.NewTaskController(mockService)
	task := models.Task{
		ID:     2,
		Name:   "task1",
		Desc:   "New task1",
		UserId: 1,
		Status: "pending",
	}

	mockService.On("Delete").Return(task, nil)
	req := httptest.NewRequest(http.MethodDelete, "/tasks/2", nil)
	rec := httptest.NewRecorder()

	c := echo.New().NewContext(req, rec)
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	err := mockController.Delete(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	//tid, _ := strconv.Atoi(id)
	mockService.AssertCalled(t, "Delete")

}

func TestFindAllTask(t *testing.T) {
	//initialize service mock
	mockService := &TaskServiceMock{}

	//create controller
	mockController := controller.NewTaskController(mockService)
	createdA := time.Now()
	layout := "2006-01-02 03:04:05 PM"
	act, err := time.Parse(layout, createdA.Format("2006-01-02 03:04:05 PM"))
	if err != nil {
		fmt.Printf("%v", err)
	}
	allTask := &[]models.Task{
		{ID: 1, Name: "task1", Desc: "New task1 Desc", CreatedAt: act, UserId: 1, Status: "pending"},
		{ID: 2, Name: "task2", Desc: "New Task2", CreatedAt: createdA, UserId: 1, Status: "completed"},
	}
	mockService.On("FindAllTask").Return(*allTask, nil)
	req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	err = mockController.FindAllTask(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	mockService.AssertCalled(t, "FindAllTask")
}
func TestFindById(t *testing.T) {
	mockService := &TaskServiceMock{}
	mockController := controller.NewTaskController(mockService)

	// Assuming CreatedAt is already in the desired format
	createdAt := time.Now()
	task := models.Task{
		ID:        1,
		Name:      "task1",
		Desc:      "New task1",
		CreatedAt: createdAt,
		UserId:    1,
		Status:    "pending",
	}
	id := strconv.Itoa(task.ID)

	// Mocking the FindByID method of the TaskService
	mockService.On("FindByID", task.ID).Return(task, nil)

	req := httptest.NewRequest(http.MethodGet, "/tasks/"+id, nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues(id)

	err := mockController.FindByID(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertCalled(t, "FindByID", task.ID)
}

func TestFindByUserID(t *testing.T) {
	mockService := &TaskServiceMock{}
	mockController := controller.NewTaskController(mockService)
	tasks := []models.Task{
		{ID: 1, Name: "task1", Desc: "New task1 Desc", UserId: 1, Status: "pending"},
		{ID: 2, Name: "task2", Desc: "New Task2", UserId: 1, Status: "completed"},
	}
	userId := 1

	// Expecting a call to FindByUserId with the provided id
	mockService.On("FindByUserId", userId).Return(tasks, nil)

	req := httptest.NewRequest(http.MethodGet, "/tasks/users/1", nil)
	rec := httptest.NewRecorder()
	//fmt.Printf("%v\n", rec)
	c := echo.New().NewContext(req, rec)
	c.SetPath("/tasks/users/:id")
	c.SetParamNames("userId")
	c.SetParamValues("1")

	err := mockController.FindByUserId(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertCalled(t, "FindByUserId", userId)
}

func TestUpdate(t *testing.T) {

	mockService := &TaskServiceMock{}
	mockController := controller.NewTaskController(mockService)

	task := &models.Task{
		ID:     1,
		Name:   "Learning unit testing",
		Desc:   "learning go test",
		UserId: 1,
		Status: "Completed",
	}

	mockService.On("Update", 1, task).Return(*task, nil)

	reqBoby := `
	{
		"desc": "learning go test",
		"id": 1,
		"name": "Learning unit testing",
		"status": "Completed",
		"user_id": 1
	}
	`
	req := httptest.NewRequest(http.MethodPut, "/tasks/1", bytes.NewBufferString(reqBoby))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := mockController.Update(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertCalled(t, "Update", 1, task)

}
