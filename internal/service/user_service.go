package service

import (
	"ToDo/internal/domain/user"
	"ToDo/internal/repository"
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
