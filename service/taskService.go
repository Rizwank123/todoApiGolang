package service

import (
	"fmt"

	"github.com/rizwank123/models"
	"github.com/rizwank123/repository"

)

type TaskService interface {
	CreateTask(task *models.Task) error
	FindAllTask() ([]models.Task, error)
	FindByID(id int) (models.Task, error)
	Delete(id int) error
	Update(id int, task *models.Task) (models.Task, error)
	FindByUserId(id int) ([]models.Task, error)
}
type taskServiceImpl struct {
	Repo  repository.TaskRepository
	Urepo repository.UserRepository
}

func NewTaskService(repo repository.TaskRepository, us repository.UserRepository) TaskService {
	return &taskServiceImpl{Repo: repo, Urepo: us}
}

func (ts *taskServiceImpl) CreateTask(task *models.Task) error {
	return ts.Repo.CreateTask(task)

}
func (ts *taskServiceImpl) FindAllTask() ([]models.Task, error) {
	return ts.Repo.FindAllTask()
}
func (ts *taskServiceImpl) FindByID(id int) (models.Task, error) {
	return ts.Repo.FindByID(id)
}
func (ts *taskServiceImpl) Delete(id int) error {
	return ts.Repo.Delete(id)
}
func (ts *taskServiceImpl) Update(id int, task *models.Task) (models.Task, error) {
	tsk, err := ts.Repo.FindByID(id)
	if err != nil {
		panic(err)
	}
	if task.Name != "" {
		tsk.Name = task.Name
	}
	if task.Desc != "" {
		tsk.Desc = task.Desc
	}
	if task.Status != tsk.Status {
		tsk.Status = task.Status
	}
	err = ts.Repo.Update(&tsk)
	return *task, err
}

func (ts *taskServiceImpl) FindByUserId(userId int) ([]models.Task, error) {
	usr, err := ts.Urepo.FindByID(userId)
	if err != nil {
		fmt.Printf("%v\n", err)

		return nil, err
	}
	if usr.Id != userId {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	task, err := ts.Repo.FindByUserId(userId)
	if err != nil {
		fmt.Printf("%v\n", err)
		//return nil, err
	}
	return task, nil
}
