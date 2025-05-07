package service

import (
	"ToDo/internal/domain/user"
	"ToDo/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetByEmail(email string) (*user.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetAll() []user.User {
	return s.repo.GetAll()
}

func (s *UserService) GetById(id uuid.UUID) (*user.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Create(u user.User) user.User {
	return s.repo.Create(u)
}

func (s *UserService) Update(id uuid.UUID, updated user.User) (*user.User, error) {
	return s.repo.Update(id, updated)
}

func (s *UserService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
