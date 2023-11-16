package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/rizwank123/models"
)

type UserRepository interface {
	FindByID(id int) (result models.User, err error)
	CreateUser(user *models.User) (usr models.User, err error)
	Update(user *models.User) (err error)
	Delete(id int) (err error)
}

type pgxUserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) (r UserRepository) {
	return &pgxUserRepository{
		db: db,
	}
}

func (r *pgxUserRepository) FindByID(id int) (result models.User, err error) {

	query := "SELECT * FROM users WHERE id=$1"

	err = r.db.QueryRow(context.Background(), query, id).Scan(&result.Id, &result.Name, &result.Email, &result.Password)

	if err != nil {

		return
	}

	return result, nil
}
func (r *pgxUserRepository) CreateUser(user *models.User) (usr models.User, err error) {
	query := "INSERT INTO users (name,email,password) values($1,$2,$3) RETURNING id"
	err = r.db.QueryRow(context.Background(), query, user.Name, user.Email, user.Password).Scan(&user.Id)

	if err != nil {
		fmt.Printf("Row not inserted%v\n", err)
	}
	fmt.Printf("user is %v\n ", usr)
	return *user, nil
}
func (r *pgxUserRepository) Update(user *models.User) (err error) {

	query := "update users set name=$1,email=$2,password=$3"
	r.db.QueryRow(context.Background(), query, user.Name, user.Email, user.Password)
	return nil
}
func (r *pgxUserRepository) Delete(id int) (err error) {
	query := "DELETE FROM users where id=$1"
	r.db.QueryRow(context.Background(), query, id)

	return nil
}
