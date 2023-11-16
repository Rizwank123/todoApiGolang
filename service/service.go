package service

import (
	"github.com/rizwank123/models"
	"github.com/rizwank123/repository"
)

type UserService interface {
	FindByID(id int) (models.User, error)
	CreateUser(user *models.User) (models.User, error)
	Update(user *models.User, id int) error
	Delete(id int) error
}

type userServiceImpl struct {
	Repo repository.UserRepository
}

func NewService(repo repository.UserRepository) UserService {
	return &userServiceImpl{Repo: repo}
}

func (s *userServiceImpl) FindByID(id int) (result models.User, err error) {
	return s.Repo.FindByID(id)
}
func (s *userServiceImpl) CreateUser(user *models.User) (models.User, error) {
	return s.Repo.CreateUser(user)
}
func (s *userServiceImpl) Update(user *models.User, id int) error {

	usr, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}
	if usr.Name != "" {
		user.Name = usr.Name

	}
	if usr.Password != "" {
		user.Password = usr.Password
	}
	if usr.Email != "" {
		user.Email = usr.Email
	}

	return s.Repo.Update(user)
}
func (s *userServiceImpl) Delete(id int) error {
	return s.Repo.Delete(id)
}
