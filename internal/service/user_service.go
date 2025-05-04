package service

import (
	"ToDo/internal/domain/user"
	"ToDo/internal/repository"
	"errors"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() []user.User {
	return s.repo.GetAll()
}

func (s *UserService) GetById(id int) (*user.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Create(u user.User) user.User {
	return s.repo.Create(u)
}

func (s *UserService) Update(id int, updated user.User) (*user.User, error) {
	return s.repo.Update(id, updated)
}

func (s *UserService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *UserService) Login(email, password string) (*user.User, error) {
	users := s.repo.GetAll()
	for _, u := range users {
		if u.Email == email && u.Password == password {
			return &u, nil
		}
	}
	return nil, errors.New("invalid email or password")
}
