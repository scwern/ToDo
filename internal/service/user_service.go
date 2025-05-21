package service

import (
	"ToDo/internal/domain/user"
	"github.com/google/uuid"
)

type UserRepositoryInterface interface {
	GetByEmail(email string) (*user.User, error)
	GetAll() ([]user.User, error)
	GetById(id uuid.UUID) (*user.User, error)
	Create(u user.User) (user.User, error)
	Update(id uuid.UUID, updated user.User) (*user.User, error)
	Delete(id uuid.UUID) error
}

type UserService struct {
	repo UserRepositoryInterface
}

func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetByEmail(email string) (*user.User, error) {
	return s.repo.GetByEmail(email)
}

func (s *UserService) GetAll() ([]user.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(id uuid.UUID) (*user.User, error) {
	return s.repo.GetById(id)
}

func (s *UserService) Create(u user.User) (user.User, error) {
	createdUser, err := s.repo.Create(u)
	if err != nil {
		return user.User{}, err
	}
	return createdUser, nil
}

func (s *UserService) Update(id uuid.UUID, updated user.User) (*user.User, error) {
	return s.repo.Update(id, updated)
}

func (s *UserService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
