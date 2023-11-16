package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rizwank123/models"

)

type TaskRepository interface {
	CreateTask(task *models.Task) error
	FindAllTask() (res []models.Task, err error)
	Update(task *models.Task) (err error)
	Delete(id int) (err error)
	FindByID(id int) (res models.Task, err error)
	FindByUserId(id int) (res []models.Task, err error)
}
type pgxTaskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) (t TaskRepository) {
	return &pgxTaskRepo{db: db}
}
func (t *pgxTaskRepo) CreateTask(task *models.Task) error {
	query := "INSERT INTO task (name, descc, created_at, user_id, status) VALUES ($1, $2, $3, $4, $5) RETURNING id	"
	timeF := time.Now().Format("02/01/2006 03:04 PM")
	err := t.db.QueryRow(context.Background(), query, task.Name, task.Desc, timeF, task.UserId, task.Status).Scan(&task.ID)
	if err != nil {
		panic(err)
	}
	// res.Name = task.Name
	// res.Desc = task.Desc
	// res.CreatedAt = task.CreatedAt
	// res.User = task.User
	return nil
}
func (t *pgxTaskRepo) FindAllTask() (res []models.Task, err error) {
	query := "SELECT * FROM task "
	rows, err := t.db.Query(context.Background(), query)

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Desc, &task.CreatedAt, &task.UserId, &task.Status); err != nil {
			fmt.Printf("%v\n", err)
		}

		res = append(res, task)
	}
	if err != nil {
		panic(err)
	}
	return res, nil

}
func (t *pgxTaskRepo) Update(task *models.Task) (err error) {
	query := "update task set name=$2,descc=$3,created_at=$4,status=$5 where id=$1"
	args := []interface{}{task.ID, task.Name, task.Desc, task.CreatedAt, task.Status}
	_, err = t.db.Exec(context.Background(), query, args...)
	fmt.Println(query, args)
	return err
}

func (t *pgxTaskRepo) Delete(id int) (err error) {
	Query := "DELETE FROM task WHERE id=$1"
	t.db.Query(context.Background(), Query, id)
	return nil
}
func (t *pgxTaskRepo) FindByID(id int) (res models.Task, err error) {
	query := "SELECT * FROM task WHERE id=$1"
	t.db.QueryRow(context.Background(), query, id).Scan(&res.ID, &res.Name, &res.Desc, &res.CreatedAt, &res.UserId, &res.Status)

	return res, nil
}
func (t *pgxTaskRepo) FindByUserId(id int) (res []models.Task, err error) {
	query := "SELECT * FROM task WHERE user_id=$1"
	rows, err := t.db.Query(context.Background(), query, id)
	if err != nil {

		return nil, err
	}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Desc, &task.CreatedAt, &task.UserId, &task.Status); err != nil {

			return nil, err
		}
		res = append(res, task)

	}
	return res, nil
}
