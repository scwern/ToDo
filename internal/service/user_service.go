package service

import (
	"ToDo/internal/domain/user"
	"github.com/google/uuid"
)

// UserRepositoryInterface определяет поведение хранилища пользователей.
type UserRepositoryInterface interface {
	GetByEmail(email string) (*user.User, error)
	GetAll() ([]user.User, error)
	GetById(id uuid.UUID) (*user.User, error)
	Create(u user.User) (user.User, error)
	Update(id uuid.UUID, updated user.User) (*user.User, error)
	Delete(id uuid.UUID) error
}

// UserService реализует бизнес-логику работы с пользователями.
type UserService struct {
	repo UserRepositoryInterface
}

// NewUserService создает новый экземпляр UserService.
func NewUserService(repo UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

// GetByEmail возвращает пользователя по email.
func (s *UserService) GetByEmail(email string) (*user.User, error) {
	return s.repo.GetByEmail(email)
}

// GetAll возвращает всех пользователей.
func (s *UserService) GetAll() ([]user.User, error) {
	return s.repo.GetAll()
}

// GetById возвращает пользователя по UUID.
func (s *UserService) GetById(id uuid.UUID) (*user.User, error) {
	return s.repo.GetById(id)
}

// Create создает нового пользователя.
func (s *UserService) Create(u user.User) (user.User, error) {
	createdUser, err := s.repo.Create(u)
	if err != nil {
		return user.User{}, err
	}
	return createdUser, nil
}

// Update обновляет существующего пользователя.
func (s *UserService) Update(id uuid.UUID, updated user.User) (*user.User, error) {
	return s.repo.Update(id, updated)
}

// Delete удаляет пользователя по UUID.
func (s *UserService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
